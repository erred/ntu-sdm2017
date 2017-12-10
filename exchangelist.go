package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"strings"
)

type PageExchange struct {
	Page
	ExchangeList
}

type ExchangeList struct {
	Requests   []ExItem
	Pending    []ExItem
	Inprogress []ExItem
	Complete   []ExItem
}

type ExItem struct {
	Id string
	U1 ExUser
	U2 ExUser
}

type ExUser struct {
	User  template.URL
	Tree  template.URL
	Skill template.URL
	State template.URL
	Name  string
}

// ==================== Handlers ====================

// /exchange/status/
func statusExchangeHandler(w http.ResponseWriter, r *http.Request) {
	paths := strings.Split(r.URL.Path, "/")
	if len(paths) != 4 {
		http.Redirect(w, r, "/exchange/status/", http.StatusFound)
		return
	}

	page, err := getPage(r)
	if errInternal(err, w) {
		log.Print("getpage err")
		return
	}
	exchange, err := getExchangeData(page.user)
	if errInternal(err, w) {
		log.Print("getex err")
		return
	}
	page.Page = "exchange"
	page.Page2 = "status"
	data := PageExchange{page, exchange}
	templates.ExecuteTemplate(w, "exchange-status.html", data)
}

// api
// /exchange/action/id/
func exchangeHandler(w http.ResponseWriter, r *http.Request) {
	paths := strings.Split(r.URL.Path, "/")
	if len(paths) != 5 {
		http.Redirect(w, r, "/exchange/status/", http.StatusFound)
	}

	user, err := getUser(r)
	if errInternal(err, w) {
		return
	}
	switch paths[2] {
	case "accept":
		err = exchangeAccept(user, paths[3])
	case "delete", "reject":
		err = exchangeDelete(user, paths[3])
	case "done":
		err = exchangeDone(user, paths[3])
	default:
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if errInternal(err, w) {
		return
	}
	w.WriteHeader(http.StatusOK)
}

// ==================== Operations ====================
func getExchangeData(user string) (ExchangeList, error) {
	var exList ExchangeList

	var myname string
	err := DB.QueryRow("SELECT name FROM user where user=?", user).Scan(&myname)
	if err != nil {
		log.Println("getExchangeData/getUser failed")
		return exList, err
	}

	rows, err := DB.Query("SELECT * FROM exchange WHERE u1=? OR u2=?", user, user)
	if err != nil {
		if err == sql.ErrNoRows {
			return exList, nil
		}
		log.Println("getExchangeData/getExchanges failed")
		return exList, err
	}
	defer rows.Close()

	for rows.Next() {
		var id string
		var u1, u2 ExUser
		err = rows.Scan(&id, &u1.User, &u1.Tree, &u1.Skill, &u1.State, &u2.User, &u2.Tree, &u2.Skill, &u2.State)
		if err != nil {
			log.Println("getExchangeData/Scan failed")
			return exList, err
		}

		it := ExItem{id, u1, u2}
		if user != u1.Name {
			it.U1 = u2
			it.U2 = u1
		}

		it.U1.Name = myname
		err = DB.QueryRow("SELECT name FROM user where name=?", it.U2.User).Scan(&it.U2.Name)
		if err != nil {
			log.Println("getExchangeData/getU2Name failed")
			return exList, err
		}

		switch it.U1.State {
		case "requested":
			exList.Requests = append(exList.Requests, it)
		case "pending":
			exList.Pending = append(exList.Pending, it)
		case "incomplete":
			exList.Inprogress = append(exList.Inprogress, it)
		case "complete":
			if it.U2.State == "incomplete" {
				exList.Inprogress = append(exList.Inprogress, it)
			} else {
				exList.Complete = append(exList.Complete, it)
			}
		default:
			log.Print("unkown state: ", it.U1.State)
		}
	}
	if err := rows.Err(); err != nil {
		log.Println("getExchangeData/Next failed")
		return exList, err
	}
	return exList, nil
}
func exchangeAccept(user string, id string) error {
	_, err := DB.Exec("UPDATE exchange SET u1state='incomplete', u2state='incomplete' WHERE exid=?", id)
	return err
}

func exchangeDone(user string, id string) error {
	var u1, u2 string
	err := DB.QueryRow("SELECT u1, u2 FROM exchange WHERE exid=?", id).Scan(&u1, &u2)
	if err != nil {
		log.Println("exchangeDone/getUser failed")
		return err
	}

	u := "u1state"
	if user == u1 {
		u = "u2state"
	}
	_, err = DB.Exec("UPDATE exchange SET "+u+"='complete' WHERE exid=?", id)
	return err
}

func exchangeDelete(user string, id string) error {
	_, err := DB.Exec("DELETE FROM exchange WHERE exid=?", id)
	return err
}

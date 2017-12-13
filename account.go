package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
)

type PageAccount struct {
	Page
	Account
}

type Account struct {
	Name  string
	Email string
}

// ==================== Handlers ====================
// shows the account page
func accountHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/account/update/" {
		err := updateAccount(r)
		if errInternal(err, w) {
			return
		}
		w.WriteHeader(http.StatusOK)
		return
	}
	if r.URL.Path != "/account/" {
		http.Redirect(w, r, "/account/", http.StatusFound)
		return
	}

	page, err := getPage(r)
	if err != nil {
		log.Println("accountHandler/getPage failed")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	account, err := getAccount(page.user)
	if err != nil {
		log.Println("accountHandler/getAccount failed")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	page.Page = "account"
	page.Page2 = "account"
	data := PageAccount{page, account}
	templates.ExecuteTemplate(w, "account.html", data)
}

// deletes user data from facebook app deauthorization
func deauthHandler(w http.ResponseWriter, r *http.Request) {
	// TODO delete user data
}

// ==================== Operations ====================

func createAccount(r *http.Request) error {
	userCk, err := r.Cookie(CkUser)
	if err != nil {
		b, _ := httputil.DumpRequest(r, false)
		log.Print(string(b))
		log.Println("createAccount/userCk failed")
		log.Println("no user id")
		return err
	}
	user := userCk.Value

	var uname string
	unameCk, err := r.Cookie(CkName)
	if err != nil {
		// b, _ := httputil.DumpRequest(r, false)
		// log.Print(string(b))
		log.Println("createAccount/unameCk failed")
		log.Println("no user name")
		// return err
		uname = ""
	} else {
		uname = unameCk.Value

	}

	var email string
	emailCk, err := r.Cookie(CkEmail)
	if err != nil {
		log.Println("createAccount/emailCk failed")
		// return err
		email = ""
	} else {
		email = emailCk.Value
	}

	_, err = DB.Exec("INSERT INTO user (user, name, email) VALUES (?, ?, ?)", user, uname, email)
	return err
}

func getAccount(user string) (Account, error) {
	var acc Account
	err := DB.QueryRow("SELECT name, email FROM user where user=?", user).Scan(&acc.Name, &acc.Email)
	return acc, err
}

func deleteAccount(user string) error {
	qs := []string{"DELETE FROM user WHERE user=?", "DELETE FROM exchange WHERE u1=? OR u2=?", "DELETE FROM tree WHERE user=?"}
	for _, q := range qs {
		_, err := DB.Exec(q, user)
		if err != nil {
			log.Println("deleteAccount/DB.Exec failed")
			return err
		}
	}
	return nil
}

type updateAccountStruct struct {
	name  string
	email string
}

func updateAccount(r *http.Request) error {
	log.Println("updating account")
	user, err := getUser(r)
	if err != nil {
		return err
	}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	var dat updateAccountStruct
	err = json.Unmarshal(b, &dat)
	if err != nil {
		return err
	}
	log.Print(string(b), dat)
	_, err = DB.Exec("UPDATE user SET name=?, email=? WHERE user=?", dat.name, dat.email, user)
	if err != nil {
		return err
	}
	log.Println("account updated")
	return nil
}

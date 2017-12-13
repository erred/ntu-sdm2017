package main

import (
	"log"
	"net/http"
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
		log.Println("createAccount/userCk failed")
		log.Println("no user id")
		return err
	}
	user := userCk.Value

	unameCk, err := r.Cookie(CkName)
	if err != nil {
		log.Println("createAccount/unameCk failed")
		log.Println("no user name")
		return err
	}
	uname := unameCk.Value

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

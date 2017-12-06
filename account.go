package main

import (
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
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	account, err := getAccount(page.user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
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
	user := userCk.Value
	if err != nil {
		return err
	}
	unameCk, err := r.Cookie(CkName)
	uname := unameCk.Value
	if err != nil {
		return err
	}
	emailCk, err := r.Cookie(CkEmail)
	email := emailCk.Value
	if err != nil {
		return err
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
			return err
		}
	}
	return nil
}

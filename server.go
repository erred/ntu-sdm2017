package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	"net/http"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

// Cookies
var CkUser = "id"
var CkName = "name"
var CkEmail = "email"
var CkLang = "lang"

// ====== Util
func errInternal(err error, w http.ResponseWriter) bool {
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return true
	}
	return false
}

func main() {
	// ==================== Flags ====================
	var port int
	flag.IntVar(&port, "port", 8080, "port to serve on ( < 65535)")
	var initDatabase bool
	flag.BoolVar(&initDatabase, "initdb", false, "(re)initialize DB")
	flag.Parse()

	// ==================== DB ====================
	err := DB.Ping()
	if err != nil {
		log.Fatal("ping db failed: ", err)
	}

	if initDatabase {
		err = initDb()
	}
	if err != nil {
		log.Fatal("init db failed: ", err)
	}

	// ==================== Tree ====================
	tree := precheck(treeHandler)
	newTree := precheck(newTreeHandler)
	updateTree := precheck(updateTreeHandler)
	deleteTree := precheck(deleteTreeHandler)
	http.Handle("/tree/", tree)
	http.Handle("/tree/new/", newTree)
	http.Handle("/tree/update/", updateTree)
	http.Handle("/tree/delete/", deleteTree)

	// ==================== Exchange ====================
	exchange := precheck(exchangeHandler)
	newExchange := precheck(newExchangeHandler)
	likeExchange := precheck(likeExchangeHandler)
	statusExchange := precheck(statusExchangeHandler)
	browseExchange := precheck(browseExchangeHandler)
	http.Handle("/exchange/", exchange)
	http.Handle("/exchange/new/", newExchange)
	http.Handle("/exchange/like/", likeExchange)
	http.Handle("/exchange/status/", statusExchange)
	http.Handle("/exchange/browse/", browseExchange)

	// ==================== Account ====================
	account := precheck(accountHandler)
	http.Handle("/account/", account)
	http.HandleFunc("/FBDeauth/", deauthHandler)
	http.HandleFunc("/account/friends/", friendsHandler)

	// ==================== Default ====================
	http.HandleFunc("/", indexHandler)

	// ==================== Static ====================
	fileserver := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
	http.Handle("/static/", fileserver)

	// ==================== Serve ====================
	fmt.Println("serving on: " + strconv.Itoa(port))
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}

// ==================== Wrapper ====================

type precheck func(w http.ResponseWriter, r *http.Request)

func (p precheck) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// ensure trailiing slash
	if !strings.HasSuffix(r.URL.Path, "/") {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	user, err := getUser(r)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusForbidden)
		return
	}

	// check user exists
	var uname string
	err = DB.QueryRow("SELECT name FROM user where user=?", user).Scan(&uname)
	if err == sql.ErrNoRows {
		err = createAccount(r)
		if err != nil {
			http.Redirect(w, r, "/", http.StatusForbidden)
			return
		}
	} else if err != nil {
		http.Redirect(w, r, "/", http.StatusForbidden)
		return
	}
	// handle good request
	p(w, r)
}

// ==================== Default ====================
func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	templates.ExecuteTemplate(w, "404.html", nil)
}

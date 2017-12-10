package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
)

// var templates, _ = template.New("").Funcs(template.FuncMap{"slice": slice}).ParseGlob("templates/*")

var templates = template.Must(template.ParseGlob("templates/*"))

func slice(n int) []int {
	return make([]int, n)
}

type Page struct {
	En      bool   // language
	user    string // userid (internal)
	Page    string //which page it is (sidebar highlight)
	Page2   string
	Sidebar []SidebarTree //sidebar
}

type SidebarTree struct {
	Name   string       //Page name
	Cname  string       // zh name
	Treeid template.URL //url to link to
}

func getPage(r *http.Request) (Page, error) {
	var page Page

	user, err := getUser(r)
	if err != nil {
		log.Println("getPage/getUser failed")
		return page, err
	}

	sidebar, err := getSidebar(user)
	if err != nil {
		log.Println("getPage/getSidebar failed")
		return page, err
	}

	page = Page{En: true, user: user, Sidebar: sidebar}
	langCk, err := r.Cookie(CkLang)
	if err == nil {
		if langCk.Value == "zh" {
			page.En = false
		}
	}
	return page, nil
}

func getUser(r *http.Request) (string, error) {
	userCk, err := r.Cookie(CkUser)
	if err != nil {
		log.Println("getUser/userCk failed")
		return "", err
	}
	return userCk.Value, nil
}

func getSidebar(user string) ([]SidebarTree, error) {
	var sidebar []SidebarTree

	rows, err := DB.Query("SELECT treeid, tree, ctree FROM treeTemplate WHERE treeid IN (SELECT treeid FROM tree WHERE user=?)", user)
	if err != nil {
		if err == sql.ErrNoRows {
			return sidebar, nil
		}
		log.Println("getSidebar/Query failed")
		return sidebar, err
	}
	defer rows.Close()

	for rows.Next() {
		var tree SidebarTree
		err = rows.Scan(&tree.Treeid, &tree.Name, &tree.Cname)
		if err != nil {
			log.Println("getSidebar/Scan failed")
			return sidebar, err
		}
		sidebar = append(sidebar, tree)
	}
	if err = rows.Err(); err != nil {
		log.Println("getSidebar/Next failed")
		return sidebar, err
	}
	return sidebar, nil
}

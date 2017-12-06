package main

import (
	"database/sql"
	"html/template"
	"net/http"
)

// var templates, _ = template.New("").Funcs(template.FuncMap{"slice": slice}).ParseGlob("templates/*")

var templates = template.Must(template.ParseGlob("templates/*"))

func slice(n int) []int {
	return make([]int, n)
}

type Page struct {
	En      bool          // language
	user    string        // userid (internal)
	Page    string        //which page it is (sidebar highlight)
	Sidebar []SidebarTree //sidebar
}

type SidebarTree struct {
	Name   string //Page name
	Cname  string // zh name
	Treeid string //url to link to
}

func getPage(r *http.Request) (Page, error) {
	var page Page

	user, err := getUser(r)
	if err != nil {
		return page, err
	}

	sidebar, err := getSidebar(user)
	if err != nil {
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
		return "", err
	}
	return userCk.Value, nil
}

func getSidebar(user string) ([]SidebarTree, error) {
	var sidebar []SidebarTree

	rows, err := DB.Query("SELECT tree, ctree FROM treeTemplate WHERE treeid IN (SELECT treeid FROM tree WHERE user=?)", user)
	if err != nil {
		if err != sql.ErrNoRows {
			return sidebar, err
		}
	}
	defer rows.Close()

	for rows.Next() {
		var tree SidebarTree
		err = rows.Scan(&tree.Name, &tree.Cname)
		if err != nil {
			return sidebar, err
		}
		sidebar = append(sidebar, tree)
	}
	if err = rows.Err(); err != nil {
		if err != sql.ErrNoRows {
			return sidebar, err
		}
	}
	return sidebar, nil
}

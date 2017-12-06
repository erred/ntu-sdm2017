package main

import (
	"net/http"
)

type IndexPage struct {
	En bool
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		notFoundHandler(w, r)
		return
	}

	var data = IndexPage{true}
	langCk, err := r.Cookie(CkLang)
	if err == nil {
		if langCk.Value == "zh" {
			data.En = false
		}
	}

	templates.ExecuteTemplate(w, "index.html", data)
}

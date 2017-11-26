package main

import "flag"
import "fmt"
import "html/template"
import "io/ioutil"
import "net/http"
import "strconv"

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "port to serve on (1024, 65535), reccommended: 49152–65535")
	flag.Parse()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	files, _ := ioutil.ReadDir("templates")
	for _, n := range files {
		name := n.Name()[:len(n.Name())-5]
		http.HandleFunc("/"+name, func(w http.ResponseWriter, r *http.Request) {
			fmt.Println(name)
			templates, _ := template.New("").Funcs(template.FuncMap{"slice": slice}).ParseGlob("templates/*")
			templates.ExecuteTemplate(w, name+".html", nil)
		})
	}

	http.HandleFunc("/", notFound)
	fmt.Println("serving on: " + strconv.Itoa(port))
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}

func slice(n int) []int {
	return make([]int, n)
}

func notFound(w http.ResponseWriter, r *http.Request) {
	templates, _ := template.New("").Funcs(template.FuncMap{"slice": slice}).ParseGlob("templates/*")
	templates.ExecuteTemplate(w, "404.html", nil)
}

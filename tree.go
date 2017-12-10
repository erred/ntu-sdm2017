package main

import (
	"database/sql"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// ==================== Template Structs
type PageTree struct {
	Page
	TreeData
}
type PageAvailTree struct {
	Page
	Trees []TreeData
}

type TreeData struct {
	Name   string
	Cname  string
	Treeid template.URL
	Svg    template.HTML
	State  string
	Liked  int
	Desc   string
	// Done map[string]bool
}

// ==================== Handlers ====================

// /tree/
// /tree/atree/
func treeHandler(w http.ResponseWriter, r *http.Request) {
	page, err := getPage(r)
	if errInternal(err, w) {
		return
	}

	paths := strings.Split(r.URL.Path, "/")
	switch len(paths) {
	case 3:
		// /tree/
		if len(page.Sidebar) == 0 {
			http.Redirect(w, r, "/tree/new/", http.StatusFound)
			return
		}
		http.Redirect(w, r, "/tree/"+string(page.Sidebar[0].Treeid)+"/", http.StatusFound)
	case 4:
		// /tree/atree/
		tree, err := getTree(page.user, paths[2])
		if err != nil {
			log.Println("treeHandler/getPage failed")
			http.Redirect(w, r, "/tree/new/"+paths[2]+"/", http.StatusFound)
			return
		}
		page.Page = "tree"
		page.Page2 = string(tree.Treeid)
		data := PageTree{page, tree}
		templates.ExecuteTemplate(w, "tree.html", data)
	default:
		http.Redirect(w, r, strings.Join(paths[:4], "/"), http.StatusFound)
	}
}

// /tree/new/
// /tree/new/atree/
func newTreeHandler(w http.ResponseWriter, r *http.Request) {
	page, err := getPage(r)
	if errInternal(err, w) {
		return
	}

	paths := strings.Split(r.URL.Path, "/")
	switch len(paths) {
	case 4:
		// /tree/new/
		trees, err := getAvailableTrees(page.user)
		if errInternal(err, w) {
			return
		}

		page.Page = "tree"
		page.Page2 = "new"
		data := PageAvailTree{page, trees}
		templates.ExecuteTemplate(w, "tree-new.html", data)
	case 5:
		// /tree/new/atree/
		err = addTree(page.user, paths[3])
		if errInternal(err, w) {
			return
		}

		http.Redirect(w, r, "/tree/"+paths[3]+"/", http.StatusFound)
	default:
		http.Redirect(w, r, strings.Join(paths[:5], ","), http.StatusFound)
	}
}

// only header reponse
// /tree/update/tree/ + POST body of state
func updateTreeHandler(w http.ResponseWriter, r *http.Request) {
	user, err := getUser(r)
	if errInternal(err, w) {
		return
	}

	paths := strings.Split(r.URL.Path, "/")
	if len(paths) != 5 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// /tree/update/atree/
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("updateTreeHandler/readall failed")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = updateTree(user, paths[3], string(b))
	if errInternal(err, w) {
		return
	}
	w.WriteHeader(http.StatusOK)
}

// /tree/delete/atree/
func deleteTreeHandler(w http.ResponseWriter, r *http.Request) {
	user, err := getUser(r)
	if errInternal(err, w) {
		return
	}

	paths := strings.Split(r.URL.Path, "/")
	if len(paths) == 5 {
		err := deleteTree(user, paths[3])
		errInternal(err, w)
	}
	http.Redirect(w, r, "/tree/", http.StatusFound)
}

// ============================ Operations ===========================
func updateTree(user string, tree string, treeState string) error {
	_, err := DB.Exec("UPDATE tree SET state=? WHERE user=? AND treeid=?", treeState, user, tree)
	return err
}
func deleteTree(user string, tree string) error {
	_, err := DB.Exec("DELETE FROM tree WHERE user=? AND treeid=?", user, tree)
	return err
}

func addTree(user string, tree string) error {
	var zeroState string
	err := DB.QueryRow("SELECT zeroState FROM treeTemplate WHERE treeid = ?", tree).Scan(&zeroState)
	if err != nil {
		log.Println("addTree/getZeroState failed")
		return err
	}

	_, err = DB.Exec("INSERT INTO tree VALUES (?, ?, ?)", user, tree, zeroState)
	return err
}

func getTree(user string, tree string) (TreeData, error) {
	var td TreeData
	td.Treeid = template.URL(tree)

	err := DB.QueryRow("SELECT state FROM tree WHERE user=? AND treeid=?", user, tree).Scan(&td.State)
	if err != nil {
		log.Println("getTree/getState failed")
		return td, err
	}

	err = DB.QueryRow("SELECT tree, ctree, svg FROM treeTemplate WHERE treeid = ?", tree).Scan(&td.Name, &td.Cname, &td.Svg)
	if err != nil {
		log.Println("getTree/getTree failed")
		return td, err
	}
	return td, nil
}

func getAvailableTrees(user string) ([]TreeData, error) {
	var treeList []TreeData

	rows, err := DB.Query("SELECT treeid, tree, ctree, svg, zeroState, desc FROM treeTemplate WHERE treeid NOT IN (SELECT treeid FROM tree WHERE user = ?)", user)
	if err != nil {
		if err == sql.ErrNoRows {
			return treeList, nil
		}
		log.Println("getAvailTree/getTree failed")
		return treeList, err
	}
	defer rows.Close()

	for rows.Next() {
		var td TreeData
		err = rows.Scan(&td.Treeid, &td.Name, &td.Cname, &td.Svg, &td.State, &td.Desc)
		if err != nil {
			log.Println("getAvailTree/Scan failed")
			return treeList, err
		}
		treeList = append(treeList, td)
	}
	if err = rows.Err(); err != nil {
		log.Println("getAvailTree/Next failed")
		return treeList, err
	}
	return treeList, nil
}

func getAllTrees(user string) ([]TreeData, error) {
	var treeList []TreeData

	rows, err := DB.Query("SELECT treeid, state FROM tree WHERE user=?", user)
	if err != nil {
		if err != sql.ErrNoRows {
			return treeList, nil
		}
		log.Println("getAllTree/Query failed")
		return treeList, err
	}
	defer rows.Close()

	for rows.Next() {
		var td TreeData
		err = rows.Scan(&td.Treeid, &td.State)
		if err != nil {
			log.Println("getAllTree/Scan failed")
			return treeList, err
		}

		err = DB.QueryRow("SELECT tree, ctree, svg FROM treeTemplate WHERE treeid=?", td.Treeid).Scan(&td.Name, &td.Cname, &td.Svg)
		if err != nil {
			log.Println("getalltree/gettree failed")
			return treeList, err
		}
		treeList = append(treeList, td)
	}
	if err = rows.Err(); err != nil {
		log.Println("getalltree/Next failed")
		return treeList, err
	}
	return treeList, nil
}

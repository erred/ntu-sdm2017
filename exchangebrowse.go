package main

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"html/template"
	"log"
	"net/http"
	"strings"
)

type PageBrowse struct {
	Page
	Items []BrowseItem
}

type BrowseItem struct {
	TreeData
	User  template.URL
	Name  string
	Liked bool
	Likes int
}

type PageNewExchange struct {
	Page
	MyTrees    []TreeData
	TargetTree TreeData
}

// ==================== Handlers ====================

func browseExchangeHandler(w http.ResponseWriter, r *http.Request) {
	paths := strings.Split(r.URL.Path, "/")
	if len(paths) != 4 {
		http.Redirect(w, r, "/exchange/browse/", http.StatusFound)
		return
	}

	page, err := getPage(r)
	if errInternal(err, w) {
		return
	}
	items, err := browseTrees(page.user)
	if errInternal(err, w) {
		return
	}
	page.Page = "exchange"
	page.Page2 = "browse"
	data := PageBrowse{page, items}
	templates.ExecuteTemplate(w, "exchange-browse.html", data)
}

// /exchange/like/user/tree/1/
func likeExchangeHandler(w http.ResponseWriter, r *http.Request) {
	paths := strings.Split(r.URL.Path, "/")
	if len(paths) != 7 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	user, err := getUser(r)
	if errInternal(err, w) {
		return
	}
	err = exchangeLike(user, paths[3], paths[4], paths[5])
	if errInternal(err, w) {
		return
	}
	w.WriteHeader(http.StatusOK)
}

func newExchangeHandler(w http.ResponseWriter, r *http.Request) {
	page, err := getPage(r)
	if errInternal(err, w) {
		return
	}

	paths := strings.Split(r.URL.Path, "/")
	switch len(paths) {
	case 5:
		// /exchange/new/user/atree/
		dataTarget, err := getTree(paths[3], paths[4])
		if errInternal(err, w) {
			return
		}
		dataMe, err := getAllTrees(page.user)
		if errInternal(err, w) {
			return
		}

		page.Page = "exchange"
		page.Page2 = "new"
		data := PageNewExchange{page, dataMe, dataTarget}
		templates.ExecuteTemplate(w, "exchange-new.html", data)
		return
	case 8:
		// /exchange/new/user/atree/askill/mytree/myskill/
		exchangeRequest(page.user, paths[6], paths[7], paths[3], paths[4], paths[5])
	}
	http.Redirect(w, r, "/exchange/status/", http.StatusFound)
}

// ==================== Operations ====================
func exchangeRequest(me string, mytree string, myskill string, user string, userTree string, userSkill string) error {
	hasher := md5.New()
	hasher.Write([]byte(me + mytree + myskill + user + userTree + userSkill))
	exid := hex.EncodeToString(hasher.Sum(nil))

	_, err := DB.Exec("INSERT INTO exchange VALUES (?, ?, ?, ?, 'requested', ?, ?, ?, 'pending')", exid, me, mytree, myskill, user, userTree, userSkill)
	return err
}

func browseTrees(user string) ([]BrowseItem, error) {
	var items []BrowseItem

	rows, err := DB.Query("SELECT user, treeid FROM tree WHERE user IS NOT ?", user)
	if err != nil {
		if err == sql.ErrNoRows {
			return items, nil
		}
		log.Println("browseTrees/DB.Query failed")
		return items, err
	}
	defer rows.Close()
	for rows.Next() {
		var bi BrowseItem
		var tid string
		err = rows.Scan(&bi.User, &tid)
		if err != nil {
			log.Println("browseTrees/rows.Scan failed")
			return items, err
		}

		bi.TreeData, err = getTree(string(bi.User), tid)
		if err != nil {
			log.Print(bi.TreeData)
			log.Println("browseTrees/getTree failed")
			return items, err
		}

		err = DB.QueryRow("SELECT name FROM user WHERE user=?", bi.User).Scan(&bi.Name)
		if err != nil {
			log.Println("browseTrees/getName failed")
			return items, err
		}

		err = DB.QueryRow("SELECT count(*) FROM likes WHERE owner=? AND treeid=?", bi.User, bi.Treeid).Scan(&bi.Likes)
		if err != nil {
			log.Println("browseTrees/countlikes failed")
			return items, err
		}

		var t string
		err = DB.QueryRow("SELECT user FROM likes WHERE user=? AND owner=? AND treeid=?", user, bi.User, bi.Treeid).Scan(&t)
		if err == sql.ErrNoRows {
			bi.Liked = false
		} else if err != nil {
			log.Println("browseTrees/ifLiked failed")
			return items, err
		} else {
			bi.Liked = true
		}
		items = append(items, bi)
	}
	if err = rows.Err(); err != nil {
		log.Println("browseTrees/rows.Next failed")
		return items, err
	}
	return items, nil
}

func exchangeLike(user string, target string, targetTree string, like string) error {
	var err error
	switch like {
	case "0":
		_, err = DB.Exec("DELETE FROM likes WHERE user=? AND owner=? AND treeid=?", user, target, targetTree)
	case "1":
		_, err = DB.Exec("INSERT INTO likes VALUES (?, ?, ?) WHERE NOT EXISTS (SELECT 1 FROM LIKES WHERE user=? AND owner=? and treeid=?)", user, target, targetTree, user, target, targetTree)
	}
	return err
}

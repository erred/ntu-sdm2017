package main

import (
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func fakeData() error {

	// Create predifined users here
	// userid , username, email
	account := [][]string{
		[]string{"0x000", "Winfred Royale", "x@example.com"},
	}
	for _, a := range account {
		_, err := DB.Exec("INSERT INTO user (user, name, email) VALUES (?, ?, ?)", a[0], a[1], a[2])
		if err != nil {
			log.Println("fakeData/account failed")
			return err
		}
	}

	// create tree templates here
	// treeid, english name, chinese name, svg, zerostate, desciption
	template := [][]string{
		[]string{tree01id, tree01name, tree01cname, tree01svg, tree01state, tree01desc},
		[]string{tree02id, tree02name, tree02cname, tree02svg, tree02state, tree02desc},
	}
	for _, t := range template {
		_, err := DB.Exec("INSERT INTO treeTemplate (treeid, tree, ctree, svg, zeroState, desc) VALUES (?, ?, ?, ?, ?, ?)", t[0], t[1], t[2], t[3], t[4], t[5])
		if err != nil {
			log.Println("fakeData/template failed")
			return err
		}
	}

	// create trees owned by user here
	// userid, treeid, state
	tree := [][]string{
		[]string{"0x000", "tree01", tree01state},
	}
	for _, t := range tree {
		_, err := DB.Exec("INSERT INTO tree (user, treeid, state) VALUES (?, ?, ?)", t[0], t[1], t[2])
		if err != nil {
			log.Println("fakeData/tree failed")
			return err
		}
	}

	// exchanges
	// hash value, (userid, treeid, skill, state) * 2
	ex := [][]string{
	// []string{"hash1", "0x00", "tree03", "laze", "requested", "0x222", "tree01", "wondernes", "pending"},
	}
	for _, e := range ex {
		_, err := DB.Exec("INSERT INTO exchange (exid, u1, u1tree, u1skill, u1state, u2, u2tree, u2skill, u2state) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)", e[0], e[1], e[2], e[3], e[4], e[5], e[6], e[7], e[8])
		if err != nil {
			log.Println("fakeData/ex failed")
			return err
		}
	}

	// likes
	// userid, treeOwnerUserId, treeId
	like := [][]string{
	// []string{"0x000", "0x111", "tree03"},
	}
	for _, l := range like {
		_, err := DB.Exec("INSERT INTO likes (user, owner, treeid) VALUES (?, ?, ?)", l[0], l[1], l[2])
		if err != nil {
			log.Println("fakeData/like failed")
			return err
		}
	}
	return nil
}

package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB, _ = sql.Open("sqlite3", "db.sqlite3")

func initDb() error {

	tables := []string{"user", "treeTemplate", "tree", "exchange", "likes"}
	dropTable := `DROP TABLE IF EXISTS `
	for _, t := range tables {
		_, err := DB.Exec(dropTable + t)
		if err != nil {
			log.Println("initDb/droptable failed: ", t)
			return err
		}
	}
	log.Println("dropped tables")

	// user is unique id
	createAccountTable := `
	CREATE TABLE user (
		user TEXT,
		name TEXT,
		email TEXT,
		friends TEXT
	)`

	// treeid is unique id
	createTemplateTable := `
	CREATE TABLE treeTemplate (
		treeid TEXT,
		tree TEXT,
		ctree TEXT,
		svg TEXT,
		zeroState TEXT,
		desc TEXT
	)`

	// state is a serialized map[string]struct{string, []string}
	createTreeTable := `
	CREATE TABLE tree (
		user TEXT,
		treeid TEXT,
		state TEXT
	)`

	// store u1 < u2
	// state:
	//	requested -> Pending
	//	complete/incomplete
	// exid is unique id
	createExchangeTable := `
	CREATE TABLE exchange (
		exid TEXT,
		u1 TEXT,
		u1tree TEXT,
		u1skill TEXT,
		u1state TEXT,
		u2 TEXT,
		u2tree TEXT,
		u2skill TEXT,
		u2state TEXT
	)`

	// user likes owner.tree
	createLikeTable := `
	CREATE TABLE likes (
		user TEXT,
		owner TEXT,
		treeid TEXT
	)`

	newTables := []string{createAccountTable, createTemplateTable, createTreeTable, createExchangeTable, createLikeTable}
	for _, t := range newTables {
		_, err := DB.Exec(t)
		if err != nil {
			log.Println("initDb/createTable failed: ", t)
			return err
		}
	}

	log.Println("created tables")
	return nil
}

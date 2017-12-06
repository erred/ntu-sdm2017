package main

import (
	"io/ioutil"
	"net/http"
)

// type FBFriend struct {
// 	Name string
// 	Id   string
// }

func friendsHandler(w http.ResponseWriter, r *http.Request) {
	user, err := getUser(r)
	if errInternal(err, w) {
		return
	}
	b, err := ioutil.ReadAll(r.Body)
	if errInternal(err, w) {
		return
	}
	err = storeFriends(user, string(b))
	if errInternal(err, w) {
		return
	}
	w.WriteHeader(http.StatusOK)
}

// func storeFriends(user string, friends []FBFriend) {
func storeFriends(user string, friends string) error {
	_, err := DB.Exec("UPDATE user SET friends=? WHERE user=?", friends, user)
	return err
}

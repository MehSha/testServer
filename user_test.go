// +build integration

package main

import (
	"net/http"
	"testing"
)

func TestAddUser(t *testing.T) {

	db := connectDB()
	defer db.Close()
	initUser(db)
	id, err := userDAM.Insert(User{
		ID:   "ID1",
		Name: "Mahdi",
	})
	if err != nil {
		t.Errorf("can not add user: %s", err)
		return
	}
	if id != "ID1" {
		t.Error("id is empty!")
		return
	}

}

func TestSVC1Integration(t *testing.T) {
	_, err := http.Get("http://127.0.0.1:8090/")
	if err != nil {
		t.Errorf("integration with SVC1 failed: %s", err)
	}
}

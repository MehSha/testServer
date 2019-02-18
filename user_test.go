// +build integration

package main

import (
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

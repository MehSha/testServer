package main

import (
	"log"

	"github.com/MehSha/basicdam"
	"github.com/jmoiron/sqlx"
)

type User struct {
	ID   string
	Name string
}

var userDAM *basicdam.BasicDAM

func initUser(db *sqlx.DB) {
	userDAM = basicdam.NewDAM(&User{}, db)
	err := userDAM.SyncDB()
	if err != nil {
		log.Fatalln(err)
	}
}

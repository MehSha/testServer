package main

import (
	"fmt"
	"time"

	"github.com/caarlos0/env"
	"github.com/jmoiron/sqlx"
)

type pgConfig struct {
	Host string `env:"POSTGRES_HOST" envDefault:"localhost"`
	DB   string `env:"POSTGRES_DB" envDefault:"test"`
	User string `env:"POSTGRES_USER" envDefault:"postgres"`
	Pass string `env:"POSTGRES_PASSWORD" envDefault:"postgres"`
}

func connectDB() *sqlx.DB {
	pgcfg := pgConfig{}
	err := env.Parse(&pgcfg)
	if err != nil {
		panic(err)
	}
	for {

		db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=disable", pgcfg.Host, pgcfg.DB, pgcfg.User, pgcfg.Pass))
		if err == nil {
			fmt.Println("DB connected successfully!")
			return db
		} else {
			fmt.Println("connecting to DB was not successful, trying again")
			time.Sleep(time.Second * 3)
		}
	}
}

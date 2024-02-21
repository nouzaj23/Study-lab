package main

import (
	"database/sql"
	"log"
	"study_lab/config"

	_ "github.com/lib/pq"
	"study_lab/api"
	db "study_lab/db/sqlc"
)

func main() {
	conf, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}
	conn, err := sql.Open(conf.DBDriver, conf.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(conf.ServerAddress)

	if err != nil {
		log.Fatal("Cannot start server" + err.Error())
	}
}

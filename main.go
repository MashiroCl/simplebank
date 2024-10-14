package main

import (
	"database/sql"
	"fmt"
	"log"
	"simplebank/api"
	db "simplebank/db/sqlc"
	"simplebank/util"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("load config failed")
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	fmt.Printf("\nDBSource is:%s\n", config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(store, config)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
	err = server.Start(config.SeverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}

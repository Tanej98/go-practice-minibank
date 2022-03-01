package main

import (
	"database/sql"
	"log"

	"github.com/Tanej98/minibank/api"
	db "github.com/Tanej98/minibank/db/sqlc"
	"github.com/Tanej98/minibank/util"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Config file not loaded:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	bank := db.NewBank(conn)
	server := api.NewServer(bank)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
}

package main

import (
	"database/sql"
	"log"

	"github.com/Wordyka/SimpleBank/api"
	"github.com/Wordyka/SimpleBank/db/util"
	db "github.com/Wordyka/SimpleBank/db/sqlc"
	_ "github.com/lib/pq"
	__ "github.com/golang/mock/mockgen/model"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}

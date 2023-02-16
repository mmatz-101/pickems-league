package api

import (
	"database/sql"
	"fmt"
	"log"

	databases "github.com/mmatz101/go-odds/databases/stores"
	"github.com/mmatz101/go-odds/utils"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var Store *databases.Store

func init() {
	config, err := utils.LoadConfig("../")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBUrl)
	if err != nil {
		log.Fatalln("unable to connect to database: ", err)
	}

	fmt.Println("Connected to database.")

	Store = databases.NewStore(conn)
}

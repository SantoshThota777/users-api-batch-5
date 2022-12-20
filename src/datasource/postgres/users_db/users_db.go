package usersdb

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	username = "postgres"
	password = "dev" // typically it should be read from environment variables
	host     = "127.0.0.1"
	schema   = "postgres"
)

var (
	Client *sql.DB
)

// init() is a built in function, which will be called when package is  imported for the first time
func init() {
	datasource := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		username,
		password,
		host,
		schema,
	)

	var err error
	Client, err = sql.Open("postgres", datasource)
	if err != nil {
		log.Fatal("Something wrong with DB!", err)
		panic(err)
	}

	if err := Client.Ping(); err != nil {
		log.Fatal("DB connection failed!", err)
	}

	fmt.Println("DB connection successful!")
}

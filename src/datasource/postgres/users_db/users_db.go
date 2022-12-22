package usersdb

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/rajesh4b8/users-api-batch-5/src/logger"
)

const (
	username = "postgres"
	password = "dev" // typically it should be read from environment variables
	host     = "127.0.0.1"
	schema   = "postgres"
)

var (
	Client *sql.DB
	log    = logger.GetLogger()
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
		log.Error("Something wrong with DB! " + err.Error())
		panic(err)
	}

	if err := Client.Ping(); err != nil {
		log.Error("DB connection failed!" + err.Error())
		panic(err)
	}

	log.Info("DB connection successful!")
}

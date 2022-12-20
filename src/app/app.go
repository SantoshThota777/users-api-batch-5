package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/rajesh4b8/users-api-batch-5/src/datasource/postgres/users_db"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {
	mapUrls()

	err := http.ListenAndServe("127.0.0.1:8080", router)
	if err != nil {
		log.Fatal("Something wrong starting the service", err)
	}
}

package app

import (
	"github.com/rajesh4b8/users-api-batch-5/src/controllers/ping"
	"github.com/rajesh4b8/users-api-batch-5/src/controllers/user"
)

func mapUrls() {
	// ping
	router.HandleFunc("/ping", ping.PingHandler).Methods("GET")

	// users
	// create user
	router.HandleFunc("/users", user.CreateUserHandler).Methods("POST")
	router.HandleFunc("/users/{userId}", user.ReadUserHandler).Methods("GET")
	router.HandleFunc("/users/{userId}", user.DeleteUserHandler).Methods("DELETE")
	router.HandleFunc("/users", user.GetAllUserHandler).Methods("GET")
}

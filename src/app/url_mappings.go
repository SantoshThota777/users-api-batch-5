package app

import (
	"github.com/rajesh4b8/users-api-batch-5/src/controllers/ping"
	"github.com/rajesh4b8/users-api-batch-5/src/controllers/user"
	"github.com/rajesh4b8/users-api-batch-5/src/services"
)

var (
	pingController = ping.NewPingController()

	userController = user.NewUserController(services.NewUserService())
)

func mapUrls() {
	// ping
	router.HandleFunc("/ping", pingController.PingHandler).Methods("GET")

	// users
	// create user
	router.HandleFunc("/users", userController.CreateUserHandler).Methods("POST")
	router.HandleFunc("/users/{userId}", userController.ReadUserHandler).Methods("GET")
}

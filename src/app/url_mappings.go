package app

import "github.com/rajesh4b8/users-api-batch-5/src/controllers/ping"

func mapUrls() {
	// ping
	router.HandleFunc("/ping", ping.PingHandler)
}

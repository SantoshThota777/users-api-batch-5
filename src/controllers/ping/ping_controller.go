package ping

import (
	"encoding/json"
	"net/http"

	"github.com/rajesh4b8/users-api-batch-5/src/logger"
)

func NewPingController() pingControllerInterface {
	return new(pingController)
}

type pingControllerInterface interface {
	PingHandler(http.ResponseWriter, *http.Request)
}

type pingController struct {
}

func (_ pingController) PingHandler(w http.ResponseWriter, r *http.Request) {
	logger.GetLogger().Info("Ping controller!")

	json.NewEncoder(w).Encode("Pong!")
}

package ping

import (
	"encoding/json"
	"net/http"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Pong!")
}

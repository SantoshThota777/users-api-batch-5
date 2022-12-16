package errors

import (
	"encoding/json"
	"net/http"
)

type RestErr struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Error   string `json:"id"`
}

func (err *RestErr) HandleError(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.Status)
	json.NewEncoder(w).Encode(err)
}

func NewBadRequestError(msg string) *RestErr {
	return &RestErr{
		Status:  http.StatusBadRequest,
		Message: msg,
		Error:   "Bad Request",
	}
}

func NewNotFoundError(msg string) *RestErr {
	return &RestErr{
		Status:  http.StatusNotFound,
		Message: msg,
		Error:   "Not Found",
	}
}

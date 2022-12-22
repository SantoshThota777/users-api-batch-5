package user

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rajesh4b8/users-api-batch-5/src/domain/users"
	"github.com/rajesh4b8/users-api-batch-5/src/services"
	"github.com/rajesh4b8/users-api-batch-5/src/utils/errors"
)

func NewUserController(userService services.UserServiceInterface) userControllerInterface {
	return userController{userService}
}

type userControllerInterface interface {
	CreateUserHandler(http.ResponseWriter, *http.Request)
	ReadUserHandler(http.ResponseWriter, *http.Request)
}

type userController struct {
	userService services.UserServiceInterface
}

func (c userController) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user users.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		// Unable to parse the request body
		errors.NewBadRequestError("Request body is not a valid json").HandleError(w)
		return
	}

	// You can have a validator
	// something like validator.Validate(user)

	u, restErr := c.userService.CreateUser(user)
	if restErr != nil {
		restErr.HandleError(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(u)
}

func (c userController) ReadUserHandler(w http.ResponseWriter, r *http.Request) {
	// get the userId from path param!
	varsMap := mux.Vars(r)

	// convert string to int
	userId, err := strconv.Atoi(varsMap["userId"])
	if err != nil {
		errors.NewBadRequestError("userId must be a number").HandleError(w)
		return
	}

	// userId is valid and it's an int
	user, restErr := c.userService.GetUser(userId)
	if restErr != nil {
		restErr.HandleError(w)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

package services

import (
	"github.com/rajesh4b8/users-api-batch-5/src/domain/users"
	"github.com/rajesh4b8/users-api-batch-5/src/utils/errors"
)

func CreateUser(u users.User) (*users.User, *errors.RestErr) {
	// Validate the input for user
	// Check if emaild is valid?
	// TODO: Validate the email id and return 400 error if not valid

	restErr := u.Save()
	if restErr != nil {
		return nil, restErr
	}

	return &u, nil
}

func GetUser(userId int) (*users.User, *errors.RestErr) {
	return users.GetUserById(userId)
}

package services

import (
	"github.com/rajesh4b8/users-api-batch-5/src/domain/users"
	"github.com/rajesh4b8/users-api-batch-5/src/utils/errors"
)

func CreateUser(u users.User) (*users.User, *errors.RestErr) {
	restErr := u.Save()
	if restErr != nil {
		return nil, restErr
	}

	return &u, nil
}

func GetUser(userId int) (*users.User, *errors.RestErr) {
	return users.GetUserById(userId)
}

func DeleteUser(userId int) *errors.RestErr {
	return users.DeleteUserById(userId)
}

func GetAllUsers() (map[int]*users.User, *errors.RestErr) {
	return users.GetAllUsers()
}

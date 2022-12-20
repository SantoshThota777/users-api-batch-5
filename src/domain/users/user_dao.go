package users

import (
	"fmt"
	"time"

	"github.com/rajesh4b8/users-api-batch-5/src/utils/errors"
)

var (
	usersDB = make(map[int]*User)
)

func (user *User) Save() *errors.RestErr {
	if usersDB[user.Id] != nil {
		return errors.NewBadRequestError(fmt.Sprintf("User with id %d already exists!", user.Id))
	}
	user.DateCreated = time.Now().String()
	usersDB[user.Id] = user

	return nil
}

func GetUserById(id int) (*User, *errors.RestErr) {
	if usersDB[id] == nil {
		return nil, errors.NewNotFoundError(fmt.Sprintf("User with id %d not found", id))
	}

	return usersDB[id], nil
}

func DeleteUserById(id int) *errors.RestErr {
	//check user exist
	if usersDB[id] == nil {
		return errors.NewNotFoundError(fmt.Sprintf("User with id %d not found", id))
	}
	// if existed delete user
	delete(usersDB, id)
	return nil
}

func GetAllUsers() ([]User, *errors.RestErr) {

	if len(usersDB) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("No Users Exist"))
	}

	users_slice := make([]User, 0, len(usersDB))
	for _, val := range usersDB {
		users_slice = append(users_slice, *val)
	}
	return users_slice, nil
}

func UpdateUserById(id int, newUserData *User) *errors.RestErr {
	//check user exist
	if usersDB[id] == nil {
		return errors.NewNotFoundError(fmt.Sprintf("User with id %d not found", id))
	}
	// if existed update user
	usersDB[id] = newUserData

	return nil
}

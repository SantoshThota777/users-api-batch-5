package users

import (
	"fmt"

	dateutils "github.com/rajesh4b8/users-api-batch-5/src/utils/date_utils"
	"github.com/rajesh4b8/users-api-batch-5/src/utils/errors"
)

var (
	usersDB = make(map[int]*User)
)

func (user *User) Save() *errors.RestErr {
	if usersDB[user.Id] != nil {
		return errors.NewBadRequestError(fmt.Sprintf("User with id %d already exists!", user.Id))
	}

	user.DateCreated = dateutils.TimeNow()
	usersDB[user.Id] = user

	return nil
}

func GetUserById(id int) (*User, *errors.RestErr) {
	if usersDB[id] == nil {
		return nil, errors.NewNotFoundError(fmt.Sprintf("User with id %d not found", id))
	}

	return usersDB[id], nil
}

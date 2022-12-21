package users

import (
	"fmt"
	"strings"

	usersdb "github.com/rajesh4b8/users-api-batch-5/src/datasource/postgres/users_db"
	"github.com/rajesh4b8/users-api-batch-5/src/utils/errors"
)

const (
	noResultsError = "no rows in result set"
	queryByUserId  = "select user_id, first_name, last_name, email_id, date_created from users where user_id = $1"
	querySaveUser  = "INSERT INTO users (first_name, last_name, email_id) values ($1, $2, $3) returning user_id, date_created"
)

var (
	usersDB = make(map[int]*User)
)

func (user *User) Save() *errors.RestErr {
	// Implementation with userDB being a in-memory map

	// if usersDB[user.Id] != nil {
	// 	return errors.NewBadRequestError(fmt.Sprintf("User with id %d already exists!", user.Id))
	// }

	// user.DateCreated = dateutils.TimeNow()
	// usersDB[user.Id] = user

	stmt, err := usersdb.Client.Prepare(querySaveUser)
	if err != nil {
		return errors.NewInternalServerError("Somethng wrong with the query! " + err.Error())
	}

	// defer will postpone executing until it's returning from the current function
	defer stmt.Close()

	result := stmt.QueryRow(user.FirstName, user.LastName, user.EmaildId)

	if err := result.Scan(&user.Id, &user.DateCreated); err != nil {
		return errors.NewInternalServerError("Somethng wrong with readng data from DB! " + err.Error())
	}

	return nil
}

func GetUserById(id int) (*User, *errors.RestErr) {
	// Implementation with userDB being a in-memory map

	// if usersDB[id] == nil {
	// 	return nil, errors.NewNotFoundError(fmt.Sprintf("User with id %d not found", id))
	// }

	// return usersDB[id], nil

	stmt, err := usersdb.Client.Prepare(queryByUserId)
	if err != nil {
		return nil, errors.NewInternalServerError("Somethng wrong with the query! " + err.Error())
	}

	// defer will postpone executing until it's returning from the current function
	defer stmt.Close()

	result := stmt.QueryRow(id)

	user := User{}

	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.EmaildId, &user.DateCreated); err != nil {
		if strings.Contains(err.Error(), noResultsError) {
			return nil, errors.NewNotFoundError(fmt.Sprintf("User with id %d not found", id))
		}

		return nil, errors.NewInternalServerError("Somethng wrong with readng data from DB! " + err.Error())
	}

	return &user, nil
}

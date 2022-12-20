package emailvalidation

import (
	"errors"
	"strings"
)

// To validate email

// Rules from google
// Usernames can contain letters (a-z), numbers (0-9), and periods (.).
// Usernames cannot contain an ampersand (&), equals sign (=), underscore (_), apostrophe ('), dash (-), plus sign (+), comma (,), brackets (<,>), or more than one period (.) in a row.
// Usernames can begin or end with non-alphanumeric characters except periods (.)
// Choose a username 6–30 characters long. Your username can be any combination of letters, numbers, or symbols.

// Constraints

// 1) should have @
// 2) cannot have  (&),(=),(_),('),(-),(+),(,),(<,>), or more than one period (.)
// 3) cannot end with (.) or start with (.)
// 4) 6–30 characters long

const (
	EMAIL_NOT_ALLOWED_CHARECTERS = "&=_'-+,><"
)

func ValidateUserEmail(email string) error {
	if !strings.ContainsAny(email, "@") {
		return errors.New("should contain @ in the username")
	}
	if strings.ContainsAny(email, EMAIL_NOT_ALLOWED_CHARECTERS) {
		return errors.New("invalid character, username cannot contain &=_'-+,><")
	}

	if len(email) < 6 || len(email) > 30 {
		return errors.New("invalid number of characters, username should be 6–30 characters long")
	}

	if string(email[0]) == "." || string(email[len(email)-1]) == "." {
		return errors.New("invalid period (.), Period cannot be at first or last character of the username")
	}
	if no_of_periods := strings.Count(email, "."); no_of_periods != 1 {
		return errors.New("username should contain only one period(.) ")
	}

	return nil

}

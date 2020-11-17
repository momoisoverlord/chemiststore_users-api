package users

import (
	"github.com/momoisoverlord/nimbuss-rest/chemiststore_users-api/utils/errors"
	"strings"
)

// User is the core of out entire microservice/API
// Define the User
type User struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

// Validate validates the user with which it is called
// It is a method on User struct
func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("Invalid Email address")
	}
	return nil
}

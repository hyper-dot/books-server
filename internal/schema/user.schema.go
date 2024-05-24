package schema

import (
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/hyper-dot/books-server/internal/utils"
)

type User struct {
	Username string `json:"username" validate:"required,min=3,max=32"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

// ValidateUser function to validate the request object
func ValidateUser(r *http.Request) (*User, error) {
	var user User

	// Decode the JSON request body into the User struct
	err := utils.ParseJSON(r.Body, &user)
	if err != nil {
		return nil, err
	}

	// Validate the User struct
	err = Validate.Struct(user)
	if err != nil {
		// Collecting validation errors
		validationErrors := err.(validator.ValidationErrors)
		return nil, validationErrors
	}

	return &user, nil
}

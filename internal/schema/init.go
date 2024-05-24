package schema

import "github.com/go-playground/validator/v10"

// Validator instance
var Validate *validator.Validate

func init() {
	Validate = validator.New()
}

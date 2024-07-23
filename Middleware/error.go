package middleware
import (
	"fmt"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

// Error represents a custom error structure
type Error struct {
	Errors map[string]interface{} `json:"errors"`
}

// NewError creates a new general error
func NewError(err error) Error {
	e := Error{}
	e.Errors = make(map[string]interface{})
	switch v := err.(type) {
	case *echo.HTTPError:
		e.Errors["body"] = v.Message
	default:
		e.Errors["body"] = v.Error()
	}
	return e
}

// NewValidatorError creates a new validation error
func NewValidatorError(err error) Error {
	e := Error{}
	e.Errors = make(map[string]interface{})
	errs := err.(validator.ValidationErrors)
	for _, v := range errs {
		e.Errors[v.Field()] = fmt.Sprintf("%v", v.Tag())
	}
	return e
}

// AccessForbidden creates a new 403 Forbidden error
func AccessForbidden() Error {
	e := Error{}
	e.Errors = make(map[string]interface{})
	e.Errors["body"] = "access forbidden"
	return e
}

// NotFound creates a new 404 Not Found error
func NotFound() Error {
	e := Error{}
	e.Errors = make(map[string]interface{})
	e.Errors["body"] = "resource not found"
	return e
}
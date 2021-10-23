package utility

import (
	"errors"
	"net/http"
	"regexp"
	"strconv"
)

func ValidateRequireAndLengthAndRegex(value string, isRequired bool, minLength, maxLength int, regex, fieldName string) error {

	length := len(value)
	Re := regexp.MustCompile(regex)
	if isRequired == true && length < 1 {
		return errors.New(fieldName + " is Required")
	}

	// Min length check
	// If params min length value is zero that indicates, there will be no min length check
	if minLength != 0 && length > 1 && length < minLength {
		return errors.New(fieldName + " must be min " + strconv.Itoa(minLength))
	}

	// Max length check
	// If params max length value is zero that indicates, there will be no max length check
	if maxLength != 0 && length > 1 && length > maxLength {
		return errors.New(fieldName + " must be max " + strconv.Itoa(maxLength))
	}

	if !Re.MatchString(value) { // Regex check
		return errors.New("invalid " + fieldName)
	}

	return nil
}

// NewHTTPError creates error model that will send as http response
// if any error occurs
func NewHTTPError(errorCode int,errors interface{}) map[string]interface{} {

	m := make(map[string]interface{})

	m["error_code"] = errorCode
	m["errors"] = errors

	return m
}

func NewValidationError(errors interface{}) map[string]interface{} {
	return NewHTTPError(http.StatusUnprocessableEntity,errors)
}

// Error codes
const (
	InvalidUserID       = "invalidUserID" // in case userid not exists
	InternalError       = "internalError" // in case, any internal server error occurs
	UserNotFound        = "userNotFound"  // if user not found
	InvalidBindingModel = "invalidBindingModel"
	EntityCreationError = "entityCreationError"
	Unauthorized        = "unauthorized" // in case, try to access restricted resource
	BadRequest          = "badRequest"
	UnprocessableEntity = "unprocessableEntity"
	UserAlreadyExists   = "userAlreadyExists"
)

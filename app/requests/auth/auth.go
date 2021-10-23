package auth

import (
	"encoding/json"
	"github.com/thearyanahmed/kloudlabllc/app/requests"
	"github.com/thedevsaddam/govalidator"
	"net/http"
	"net/url"
)

type (
	RegisterUser struct {
		Name string
		Email string
		Password string //`json:"password"`

		requests.RequestError
	}
)

func RegisterUserRequest(r *http.Request) *RegisterUser {
	rules := govalidator.MapData{
		"name": []string{"required", "between:3,12"},
		"email":    []string{"required", "email"},
		"password":    []string{"required", "between:4,40"},
	}

	messages := govalidator.MapData{
		"name": []string{"required:name is required", "between:name must be between 3 to 12 characters"},
		"email":    []string{"required:email is required", "email:email must of valid format"},
		"password":    []string{"required:password is required", "between:password must be between 4 to 40 characters"},
	}

	opts := govalidator.Options{
		Request:         r,        // request object
		Rules:           rules,    // rules map
		Messages:        messages, // custom message map (Optional)
		RequiredDefault: true,     // all the field to be pass the rules
	}

	v := govalidator.New(opts)

	var registerUser RegisterUser

	err := json.NewDecoder(r.Body).Decode(&registerUser)

	if err != nil {
		decodeErr := url.Values{}
		decodeErr.Set("decode_error",err.Error())

		registerUser.ErrorBag = decodeErr
	}

	e := v.Validate()
	registerUser.ErrorBag = e

	return &registerUser
}
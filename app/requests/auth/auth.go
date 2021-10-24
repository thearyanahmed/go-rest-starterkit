package auth

import (
	"fmt"
	_ "fmt"
	_ "github.com/thearyanahmed/kloudlabllc/app/requests"
	"github.com/thedevsaddam/govalidator"
	"net/http"
	"net/url"
)

type (
	RegisterUser struct {
		Name string `json:"name"`
		Email string `json:"email"`
		Password string `json:"password"`

		//requests.RequestError
	}
)

func RegisterUserRequest(r *http.Request) (*RegisterUser,url.Values) {
	var registerUser RegisterUser

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
		Data: &registerUser,
		Request:         r,        // request object
		Rules:           rules,    // rules map
		Messages:        messages, // custom message map (Optional)
		RequiredDefault: true,     // all the field to be pass the rules
	}

	v := govalidator.New(opts)

	fmt.Println("reg user->",registerUser,"reg iser email",registerUser.Email)

	return  &registerUser, v.ValidateJSON()
}
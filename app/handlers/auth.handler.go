package handlers

import (
	"encoding/json"
	"golang-mongodb-restful-starter-kit/app/models"
	"golang-mongodb-restful-starter-kit/app/services/auth"
	"golang-mongodb-restful-starter-kit/app/services/jwt"
	"golang-mongodb-restful-starter-kit/config"
	"golang-mongodb-restful-starter-kit/utility"
	"net/http"
)

// AuthHandler ..
type AuthHandler struct {
	service auth.AuthServiceInterface
	conf    *config.Configuration
}

func NewAuthAPI(authSrv auth.AuthServiceInterface, conf *config.Configuration) *AuthHandler {
	return &AuthHandler{
		service: authSrv,
		conf:    conf,
	}
}

func (h *AuthHandler) Create(w http.ResponseWriter, r *http.Request) {
	payload := new(signupReq)
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&payload)

	requestUser := &models.User{Name: payload.Name, Email: payload.Email, Password: payload.Password}
	result := make(map[string]interface{})

	if validateError := requestUser.Validate(); validateError != nil {
		result = utility.NewHTTPCustomError(utility.BadRequest, validateError.Error())
		utility.Response(w, result,http.StatusUnprocessableEntity)
		return
	}

	requestUser.Initialize()

	if h.service.IsUserAlreadyExists(r.Context(), requestUser.Email) {
		result = utility.NewHTTPError(utility.UserAlreadyExists)
		utility.Response(w, result,http.StatusUnprocessableEntity)
		return
	}

	err := h.service.Create(r.Context(), requestUser)

	if err != nil {
		result = utility.NewHTTPError(utility.EntityCreationError)
		utility.Response(w, result,http.StatusBadRequest)
		return
	}

	result = utility.SuccessPayload(nil, "Successfully registered")
	utility.Response(w, result,http.StatusCreated)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	credentials := new(models.Credential)
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&credentials)

	user, err := h.service.Login(r.Context(), credentials)

	if err != nil || user == nil {

		result := utility.NewHTTPError(utility.Unauthorized)
		utility.Response(w, result,http.StatusUnprocessableEntity)
		return
	}

	j := jwt.JwtToken{C: h.conf}

	tokenMap, err := j.CreateToken(user.ID.Hex(), user.Role)

	if err != nil {
		result := utility.NewHTTPError(utility.InternalError)
		utility.Response(w, result,http.StatusInternalServerError)
		return
	}

	res := &loginRes{
		Token: tokenMap["token"],
		User:  user,
	}

	result := utility.SuccessPayload(res, "successfully logged In")
	utility.Response(w, result,http.StatusOK)
}

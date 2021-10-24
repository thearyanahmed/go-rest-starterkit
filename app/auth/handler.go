package auth

import (
	"encoding/json"
	"fmt"
	"github.com/thearyanahmed/kloudlabllc/app/services/jwt"
	"github.com/thearyanahmed/kloudlabllc/app/user"
	"github.com/thearyanahmed/kloudlabllc/config"
	"github.com/thearyanahmed/kloudlabllc/utility"
	"net/http"
)

type Handler struct {
	service ServiceInterface
	conf    *config.Configuration
}

func NewAuthAPI(authSrv ServiceInterface, conf *config.Configuration) *Handler {
	return &Handler{
		service: authSrv,
		conf:    conf,
	}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	//payload := new(signupReq)
	//defer r.Body.Close()
	//
	//decoder := json.NewDecoder(r.Body)
	//decoder.Decode(&payload)
	//
	//fmt.Println("1 payload",payload)

	//requestUser := &models.User{Name: payload.Name, Email: payload.Email, Password: payload.Password}
	result := make(map[string]interface{})

	userData, err := registerUserRequest(r)

	fmt.Println(userData)

	if len(err) > 0 {
		result = utility.NewValidationError(err)
		utility.Response(w, result,http.StatusUnprocessableEntity)

		return
	} else {
		utility.Response(w, userData,http.StatusOK)

		return
	}

	//if len(e) > 0 {
	//	result = utility.NewValidationError(e)
	//	utility.Response(w, result,http.StatusUnprocessableEntity)
	//
	//	return
	//} else {
	//
	//	utility.Response(w, payload,http.StatusOK)
	//
	//	return
	//}

	//
	//if validateError := requestUser.Validate(); validateError != nil {
	//	result = utility.NewHTTPError(utility.BadRequest, validateError.Error())
	//}


	//break
	//requestUser.Initialize()
	//
	//if h.service.IsUserAlreadyExists(r.Context(), requestUser.Email) {
	//	result = utility.NewHTTPError(http.StatusUnprocessableEntity,nil)
	//	utility.Response(w, result,http.StatusUnprocessableEntity)
	//	return
	//}
	//
	//err := h.service.Create(r.Context(), requestUser)
	//
	//if err != nil {
	//	result = utility.NewHTTPError(http.StatusUnprocessableEntity,nil)
	//	utility.Response(w, result,http.StatusBadRequest)
	//	return
	//}
	//
	//result = utility.SuccessPayload(nil, "successfully registered.")
	//utility.Response(w, result,http.StatusCreated)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	credentials := new(user.Credential)
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&credentials)

	user, err := h.service.Login(r.Context(), credentials)

	if err != nil || user == nil {


		result := utility.NewHTTPError(http.StatusUnauthorized,nil)
		utility.Response(w, result,http.StatusUnprocessableEntity)
		return
	}

	j := jwt.JwtToken{C: h.conf}

	tokenMap, err := j.CreateToken(user.ID.Hex(), user.Role)

	if err != nil {
		result := utility.NewHTTPError(http.StatusUnauthorized,nil)
		utility.Response(w, result,http.StatusInternalServerError)
		return
	}

	res := struct {
		Message string
		TokenMap map[string]string
	}{
		Message: "Implement me",
		TokenMap: tokenMap,
	}

	result := utility.SuccessPayload(res, "successfully logged in")
	utility.Response(w, result,http.StatusOK)
}

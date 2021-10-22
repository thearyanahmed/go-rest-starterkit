package handlers

import (
	"encoding/json"
	"golang-mongodb-restful-starter-kit/app/models"
	userSrv "golang-mongodb-restful-starter-kit/app/services/user"
	"golang-mongodb-restful-starter-kit/utility"
	"net/http"
)

// UserHandler - handles user request
type UserHandler struct {
	service userSrv.UserServiceInterface
}

func NewUserAPI(userService userSrv.UserServiceInterface) *UserHandler {
	return &UserHandler{
		service: userService,
	}
}

func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	user, err := h.service.Get(r.Context(), utility.GetLoggedInUserID(r))

	if err != nil {
		utility.Response(w, utility.NewHTTPError(utility.InternalError),http.StatusInternalServerError)
		return
	}

	utility.Response(w, utility.SuccessPayload(user, ""),http.StatusOK)
}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	updatedUser := new(models.UserUpdate)

	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&updatedUser)

	result := make(map[string]interface{})
	err := h.service.Update(r.Context(), utility.GetLoggedInUserID(r), updatedUser)

	if err != nil {
		result = utility.NewHTTPCustomError(utility.BadRequest, err.Error())
		utility.Response(w, result,http.StatusUnprocessableEntity)
		return
	}

	result = utility.SuccessPayload(nil, "successfully updated.")
	utility.Response(w, result,http.StatusOK)

}

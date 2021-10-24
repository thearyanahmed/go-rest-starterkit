package user

import (
	"encoding/json"
	"github.com/thearyanahmed/kloudlabllc/utility"
	"net/http"
)

type Handler struct {
	service ServiceInterface
}

func NewUserAPI(userService ServiceInterface) *Handler {
	return &Handler{
		service: userService,
	}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	user, err := h.service.Get(r.Context(), utility.GetLoggedInUserID(r))

	if err != nil {
		utility.Response(w, utility.NewHTTPError(http.StatusServiceUnavailable,nil),http.StatusInternalServerError)
		return
	}

	utility.Response(w, utility.SuccessPayload(user, ""),http.StatusOK)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	updatedUser := new(UserUpdate)

	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&updatedUser)

	result := make(map[string]interface{})
	err := h.service.Update(r.Context(), utility.GetLoggedInUserID(r), updatedUser)

	if err != nil {
		result = utility.NewHTTPError(http.StatusBadRequest, err.Error())
		utility.Response(w, result,http.StatusUnprocessableEntity)
		return
	}

	result = utility.SuccessPayload(nil, "successfully updated.")
	utility.Response(w, result,http.StatusOK)

}


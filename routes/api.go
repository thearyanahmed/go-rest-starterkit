package routes

import (
	"github.com/thearyanahmed/kloudlabllc/app/auth"
	"github.com/thearyanahmed/kloudlabllc/app/user"
	"github.com/thearyanahmed/kloudlabllc/config"
	"net/http"

	"github.com/gorilla/mux"
	mgo "gopkg.in/mgo.v2"
)

var (
	BaseRoute = "/api/v1"
)

func InitializeRoutes(router *mux.Router, dbSession *mgo.Session, conf *config.Configuration) {
	userRepository := user.NewRepository(dbSession, conf)
	userService := user.NewService(userRepository)
	userAPI := user.NewUserAPI(userService)

	authService := auth.NewService(userRepository)

	authAPI := auth.NewAuthAPI(authService, conf)

	apiV1 := router.PathPrefix(BaseRoute).Subrouter()
	//  -------------------------- Auth APIs ------------------------------------
	apiV1.HandleFunc("/auth/register", authAPI.Create).Methods(http.MethodPost)
	apiV1.HandleFunc("/auth/login", authAPI.Login).Methods(http.MethodPost)

	// -------------------------- User APIs ------------------------------------
	apiV1.HandleFunc("/users/me", userAPI.Get).Methods(http.MethodGet)
	apiV1.HandleFunc("/users", userAPI.Update).Methods(http.MethodPatch)
}

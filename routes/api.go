package routes

import (
	api "github.com/thearyanahmed/kloudlabllc/app/handlers"
	userRepo "github.com/thearyanahmed/kloudlabllc/app/repositories/user"
	authSrv "github.com/thearyanahmed/kloudlabllc/app/services/auth"
	userSrv "github.com/thearyanahmed/kloudlabllc/app/services/user"
	"github.com/thearyanahmed/kloudlabllc/config"
	"net/http"

	"github.com/gorilla/mux"
	mgo "gopkg.in/mgo.v2"
)

var (
	BaseRoute = "/api/v1"
)

func InitializeRoutes(router *mux.Router, dbSession *mgo.Session, conf *config.Configuration) {
	userRepository := userRepo.New(dbSession, conf)
	userService := userSrv.New(userRepository)
	authService := authSrv.New(userRepository)
	authAPI := api.NewAuthAPI(authService, conf)
	userAPI := api.NewUserAPI(userService)

	// Routes

	//  -------------------------- Auth APIs ------------------------------------
	router.HandleFunc(BaseRoute+"/auth/register", authAPI.Create).Methods(http.MethodPost)
	router.HandleFunc(BaseRoute+"/auth/login", authAPI.Login).Methods(http.MethodPost)

	// -------------------------- User APIs ------------------------------------
	router.HandleFunc(BaseRoute+"/users/me", userAPI.Get).Methods(http.MethodGet)
	router.HandleFunc(BaseRoute+"/users", userAPI.Update).Methods(http.MethodPatch)
}

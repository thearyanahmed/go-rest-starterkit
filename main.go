package main

import (
	"golang-mongodb-restful-starter-kit/app/middleware"
	"golang-mongodb-restful-starter-kit/app/services/jwt"
	"golang-mongodb-restful-starter-kit/config"
	"golang-mongodb-restful-starter-kit/db"
	_ "golang-mongodb-restful-starter-kit/docs"
	"golang-mongodb-restful-starter-kit/routes"
	"golang-mongodb-restful-starter-kit/utility"

	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	mgo "gopkg.in/mgo.v2"
)

func main() {

	// Initialize config
	conf := config.NewConfig()

	// Make connection with db and get instance
	dbSession := db.GetInstance(conf)

	//
	dbSession.SetSafe(&mgo.Safe{})

	// Router
	router := mux.NewRouter()
	routes.InitializeRoutes(router, dbSession, conf)
	// JWT services
	jwtService := jwt.JwtToken{C: conf}

	// Added middleware over all request to authenticate
	router.Use(middleware.Cors, jwtService.ProtectedEndpoint)

	// Server configuration
	srv := &http.Server{
		Handler:      utility.Headers(router), // Set header to routes
		Addr:         conf.Address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Application is running at ", conf.Address)

	// Serving application at specified port
	log.Fatal(srv.ListenAndServe())

}

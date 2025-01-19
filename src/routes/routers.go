package routes

import (
	"backend_CMS_Golang/src/controllers"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	// User Routes
	router.HandleFunc("/api/user", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/api/user/{id}", controllers.GetUser).Methods("GET")
	router.HandleFunc("/api/user", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/api/user/{id}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/user/{id}", controllers.DeleteUser).Methods("DELETE")

	// Auth Routes
	router.HandleFunc("/auth/login", controllers.LoginWithPassword).Methods("POST")
	router.HandleFunc("/auth/google-login", controllers.LoginWtihGoogle).Methods("POST")

	return router
}

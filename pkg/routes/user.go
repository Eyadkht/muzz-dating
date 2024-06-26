package routes

import (
	"dating-app/pkg/handlers"
	"net/http"
)

func RegisterUserRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/login", handlers.UserLogin)
	mux.HandleFunc("/user/create", handlers.CreateUser)
}

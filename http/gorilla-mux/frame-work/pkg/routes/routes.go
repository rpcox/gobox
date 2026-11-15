package routes

import (
	"github.com/gorilla/mux"
	"github.com/rpcox/gobox/http/gorilla-mux/frame-work/pkg/handlers"
	"github.com/rpcox/gobox/http/gorilla-mux/frame-work/pkg/middleware"
)

func RegisterRoutes(r *mux.Router) {
	r.Use(middleware.Recovery)
	r.Use(middleware.Logging)
	r.Use(middleware.Auth)

	r.HandleFunc("/", handlers.HomeHandler).Methods("GET")

	userRouter := r.PathPrefix("/users").Subrouter()
	userRouter.HandleFunc("/", handlers.GetAllUsers).Methods("GET")
	userRouter.HandleFunc("/{id}", handlers.GetUserById).Methods("GET")

	panicRouter := r.PathPrefix("/panic").Subrouter()
	panicRouter.HandleFunc("/{bool}", handlers.GetPanic).Methods("GET")
}

// SDG

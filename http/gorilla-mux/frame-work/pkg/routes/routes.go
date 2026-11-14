package routes

import (
    "github.com/gorilla/mux"
    "github.com/rpcox/gobox/http/gorilla-mux/framework/pkg/handlers"
    "github.com/rpcox/gobox/http/gorilla-mux/framework/pkg/middleware"
)

func RegisterRoutes(r *mux.Router) {
    r.Use(middleware.LoggingMiddleware)

    r.HandleFunc("/", handlers.HomeHandler).Methods("GET")
    r.HandleFunc("/users/{id}", handlers.UserHandler).Methods("GET")
}


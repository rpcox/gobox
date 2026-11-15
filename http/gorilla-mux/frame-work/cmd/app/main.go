package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rpcox/gobox/http/gorilla-mux/frame-work/pkg/routes"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterRoutes(router)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Printf("listening @ %s\n", server.Addr)
	log.Fatal(server.ListenAndServe())
}

// SDG

package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Response struct {
	Message string `json:"message"`
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("HomeHandler")
	response := Response{Message: "hello world basic api"}
	json.NewEncoder(w).Encode(response)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	log.Println("GetAllUsers")
	response := Response{Message: "list=user1,user2,user3"}
	json.NewEncoder(w).Encode(response)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	log.Println("GetUserById")
	vars := mux.Vars(r)
	response := Response{Message: "user_id=" + vars["id"]}
	json.NewEncoder(w).Encode(response)
}

func GetPanic(w http.ResponseWriter, r *http.Request) {
	log.Println("GetPanic")
	var response Response
	vars := mux.Vars(r)
	switch vars["bool"] {
	case "true":
		response = Response{Message: "internal server error test"}
		json.NewEncoder(w).Encode(response)
		panic(`throwing panic`)
		//json.NewEncoder(w).Encode(response)
	case "false":
		response = Response{Message: "panic == false"}
	default:
		response = Response{Message: "panic == unknown"}
	}

	json.NewEncoder(w).Encode(response)
}

// SDG

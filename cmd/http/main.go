package main

import (
	"fmt"
	"github.com/abgsatman/mini-go-framework/internal/http/userapi"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	repo := userapi.NewInMemoryUserRepository()
	service := userapi.NewUserService(repo)
	handler := userapi.NewUserHandler(service)

	router.HandleFunc("/getallusers", handler.GetAllUsers).Methods("GET")
	router.HandleFunc("/getuserbyid/{id}", handler.GetUserById).Methods("GET")
	router.HandleFunc("/createuser", handler.CreateUser).Methods("POST")

	fmt.Println("Server is running at localhost:8080")
	http.ListenAndServe(":8080", router)
}

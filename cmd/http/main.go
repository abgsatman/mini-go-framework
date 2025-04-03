package main

import (
	"fmt"
	"net/http"

	"github.com/abgsatman/mini-go-framework/internal/http/userapi"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hc", userapi.HcHandler).Methods("GET")
	r.HandleFunc("/getuserbyid/{id}", userapi.GetUserByIdHandler).Methods("GET")

	fmt.Println("Server is running at localhost:8080")
	http.ListenAndServe(":8080", r)

}

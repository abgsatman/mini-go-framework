package main

import (
	"fmt"

	"github.com/abgsatman/minigoapi/internal/handler"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hc", handler.HcHandler).Methods("GET")

	fmt.Println("Server is running at localhost:8080")
}

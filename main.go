package main

import (
	"fmt"
	"net/http"

	"github.com/abgsatman/minigoapi/internal/handler"
)

func main() {
	fmt.Println("Mini GO API")

	http.HandleFunc("/hc", handler.HcHandler)
	http.ListenAndServe(":8080", nil)

	fmt.Println("Server is running at localhost:8080")
}

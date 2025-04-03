package handler

import (
	"fmt"
	"net/http"
)

func HcHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("working well..."))

	fmt.Println("working well...")
}

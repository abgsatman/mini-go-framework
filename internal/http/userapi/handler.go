package userapi

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func HcHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("working well..."))

	fmt.Println("working well...")
}

func GetUserByIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprintf(w, "id: %s", id)

	fmt.Println("id", id)

}

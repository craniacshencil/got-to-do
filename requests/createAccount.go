package requests

import (
	"log"
	"net/http"
)

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	log.Println("Control actually entered here")
}

package requests

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	responses "github.com/craniacshencil/got_to_do/responses"
)

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	res, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Reading response ERR:", err)
	}

	var signupData responses.SignupForm
	err = json.Unmarshal(res, &signupData)
	if err != nil {
		log.Fatal("Unmarshalling json ERR:", err)
	}
}

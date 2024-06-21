package requests

import (
	"fmt"
	"net/http"
)

func CreateAccount(w http.ResponseWriter, r *http.Request){
    r.ParseForm()
    fmt.Println("control entered here")
    fmt.Println(r.Form)
}

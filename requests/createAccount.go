package requests

import (
	"fmt"
	"log"
	"net/http"
)

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)

	password := r.FormValue("password")
	confirmPassword := r.FormValue("confirmPassword")

	// password validation
	isPasswordValid, err := isPasswordStrong(password)
	if err != nil {
		log.Println("ERROR:", err)
	}
	log.Println(isPasswordValid)

	// password, confirm-password match
	if password == confirmPassword {
		log.Println("Passwords match")
	} else {
		log.Println("Passwords don't match")
	}
}

func isPasswordStrong(password string) (bool, error) {
	if len(password) < 8 {
		return false, fmt.Errorf("Password smaller than 8 characters")
	}

	var capitalPresent, specialCharPresent, numberPresent bool
	for _, charAscii := range password {
		if (charAscii < 91) && (charAscii > 64) {
			fmt.Printf("capital leter: %c:%v\n", charAscii, charAscii)
			capitalPresent = true
		} else if (charAscii > 47) && (charAscii < 58) {
			fmt.Printf("number: %c:%v\n", charAscii, charAscii)
			numberPresent = true
		} else if (charAscii < 97) || (charAscii > 122) {
			fmt.Printf("special character: %c:%v\n", charAscii, charAscii)
			specialCharPresent = true
		}
	}

	if !capitalPresent {
		return false, fmt.Errorf("Capital letter not present")
	} else if !numberPresent {
		return false, fmt.Errorf("Number not present")
	} else if !specialCharPresent {
		return false, fmt.Errorf("Special character not present")
	}

	return true, nil
}

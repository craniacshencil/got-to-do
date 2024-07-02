package validation

import (
	"fmt"
)

func ValidateSignup(password, confirmPassword string) error {
	// password validation
	err := isPasswordStrong(password)
	if err != nil {
		return err
	}

	// password, confirm-password don't match
	if password != confirmPassword {
		return fmt.Errorf("passwords don't match")
	}

	// when password is strong and password, confirm-password match
	return nil
}

func isPasswordStrong(password string) error {
	if len(password) < 8 {
		return fmt.Errorf("password smaller than 8 characters")
	}

	var capitalPresent, specialCharPresent, numberPresent bool
	for _, charAscii := range password {
		if (charAscii < 91) && (charAscii > 64) {
			capitalPresent = true
		} else if (charAscii > 47) && (charAscii < 58) {
			numberPresent = true
		} else if (charAscii < 97) || (charAscii > 122) {
			specialCharPresent = true
		}
	}

	if !capitalPresent {
		return fmt.Errorf("capital letter not present")
	} else if !numberPresent {
		return fmt.Errorf("number not present")
	} else if !specialCharPresent {
		return fmt.Errorf("special character not present")
	}

	return nil
}

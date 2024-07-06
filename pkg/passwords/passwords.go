package passwords

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Err: Password couldn't be hashed:", err)
		return nil, err
	}
	return hash, nil
}

func MatchPassword(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		log.Println("Err: Password couldn't be hashed:", err)
		return err
	}
	return nil
}

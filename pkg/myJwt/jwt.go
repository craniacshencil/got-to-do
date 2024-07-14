package myJwt

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func getSecretKey() ([]byte, error) {
	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Println("ERR: While loading .env file")
		return nil, err
	}

	// Get secret string from .env
	jwtSecretString := os.Getenv("SECRET_JWT_KEY")
	if jwtSecretString == "" {
		log.Println("ERR: Couldn't find jwt secret key in .env")
		return nil, err

	}
	jwtSecretKey := []byte(jwtSecretString)
	return jwtSecretKey, nil
}

func CreateToken(userID string) (string, error) {
	// Creating token with the claims
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"iss": "got-to-do",
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour).Unix(),
	})

	// Getting the secret key
	secretKey, err := getSecretKey()
	if err != nil {
		log.Println(err)
		return "", err
	}

	// Signing the key with HMAC SHA-256
	tokenString, err := claims.SignedString(secretKey)
	if err != nil {
		log.Println("ERR: While signing the string,", err)
		return "", err
	}
	return tokenString, nil
}

func validateToken(tokenString string) (*jwt.Token, error) {
	// parsing the tokenString to create an actual jwt Token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Println("ERR: Signing method was not HMAC-SHA256 it is:", token.Header["alg"])
			return nil, fmt.Errorf(
				"ERR: Signing method was not HMAC-SHA-256 it is: %v",
				token.Header["alg"],
			)
		}
		return getSecretKey()
	})
	// Case when there is an error while parsing the JSON
	if err != nil {
		log.Println("ERR:", err)
		return nil, err
	}

	// Case when the token is not valid
	if !token.Valid {
		log.Println("ERR: The token is not valid")
		return nil, fmt.Errorf("ERR: The token is not valid")
	}

	return token, nil
}

func AuthorizeUser(r *http.Request) (*jwt.Token, error) {
	cookie, err := r.Cookie("jwt")
	if err != nil {
		log.Println("ERR: Couldn't find cookie", err)
		return nil, fmt.Errorf("couldn't find cookie: %v", err)
	}

	// Validate the cookie, store userID
	token, err := validateToken(cookie.Value)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("%v", err)
	}

	return token, nil
}

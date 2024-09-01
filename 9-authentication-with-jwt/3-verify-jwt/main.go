package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("सीेेकरेटःकी")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func verifyJWT(tokenString string) (*Claims, error) {
	claims := Claims{}

	// Parse the token using the same secret key and extract the claims
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, fmt.Errorf("invalid token signature")
		}
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return &claims, nil
}

func main() {
	// Simulate receiving this token and verifying it
	claims, err := verifyJWT("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Im5ld3VzZXIiLCJleHAiOjE3MjUxNjcwNjR9.-LmKhNFIcyXNOZLTP1uoSkh13JLgLp2WFX6VaR3hPAU")
	if err != nil {
		fmt.Println("Error verifying token:", err)
		return
	}

	// Output the username stored in the token
	fmt.Println("Token verified! Username:", claims.Username)
}

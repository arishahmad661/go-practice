// JWT is a widely-used standard for securely transmitting information between parties as a JSON object.
// In the context of backend development, JWTs are often used to secure APIs by managing user authentication.

// How JWT Works
// 1. User logs in: The server verifies the user’s credentials and, if valid, creates a JWT.
// 2. JWT creation: The JWT typically includes a payload (user data), a header (metadata), and a signature (for verifying the token).
// 3. Client stores JWT: The client stores the JWT (usually in local storage or a cookie) and sends it in the Authorization header for subsequent requests.
// 4. Server verifies JWT: For each protected endpoint, the server verifies the JWT. If valid, the server processes the request.

// Three Parts of a JWT:
// Header: Contains metadata about the token, such as the algorithm used for signing.
// Payload: Contains the claims, which are statements about the user or other data.
// Signature: Used to verify the authenticity of the token.

package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("सीेेकरेटःकी")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func generateJWT(username string) (string, error) {
	expirationTime := time.Now().Add(15 * time.Minute)
	claims := &Claims{
		Username:       username,
		StandardClaims: jwt.StandardClaims{ExpiresAt: expirationTime.Unix()},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	tokenString, err := generateJWT("newuser")
	if err != nil {
		panic(err)
	}
	fmt.Println(tokenString)
}

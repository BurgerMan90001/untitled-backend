package auth

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func JWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

func CreateJWTToken(userID string) (string, error) {

	signingKey := getJWTSecretKey()
	// expires in 24 hours
	//expirationTime := time.Now().Add(24 * time.Hour)

	// expires in 3 minutes test
	expirationTime := time.Now().Add(3 * time.Minute)

	//claims := jwt.StandardClaims{ExpiresAt: expirationTime.Unix(), Id: userID}
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expirationTime),
		Issuer: "test",
		
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(signingKey)

	return tokenString, err
}

func VerifyJWTToken(tokenString string) {
	claims := &jwt.RegisteredClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, keyFunc, jwt.WithLeeway(5*time.Second))

	if err != nil {
		log.Fatal(err)
	} else if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok {
		fmt.Println(claims.ExpiresAt, claims.Issuer)
	} else {
		log.Fatal("unknown claims type, cannot proceed")
	}

}

func getJWTSecretKey() []byte {
	return []byte(os.Getenv("JWT_SECRET"))
}

func keyFunc(token *jwt.Token) (any, error) {
	return getJWTSecretKey(), nil
}

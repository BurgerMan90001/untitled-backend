package auth

import (
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

	// expires in 24 hours
	//expirationTime := time.Now().Add(24 * time.Hour)

	// expires in 3 minutes test
	expirationTime := time.Now().Add(3 * time.Minute)

	//claims := jwt.StandardClaims{ExpiresAt: expirationTime.Unix(), Id: userID}
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = expirationTime
	claims["user"] = "adasdasdafaegdf"

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	return tokenString, err
}

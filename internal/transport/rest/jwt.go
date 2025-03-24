package rest

import (
	"indication/internal/transport/rest/helpers"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

type Credentials struct {
	Username string `json:"username"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var SecretKey = []byte("secretkey")

func generateToken(username string) (string, error) {
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(90 * time.Second).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func validateToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			//http.Error(w, "Missing authorization token", http.StatusUnauthorized)
			helpers.ReturnResonse(w, "Missing authorization token", http.StatusUnauthorized)
			return
		}
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return SecretKey, nil
		})
		if err != nil || !token.Valid {
			helpers.ReturnResonse(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r) // Call the next handler
	})
}

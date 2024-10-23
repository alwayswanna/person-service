package handlers

import (
	"crypto/rsa"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
)

const ProtectedPattern = "/api/v1"

var rsaKey *rsa.PublicKey

func Init(key *rsa.PublicKey) {
	rsaKey = key
}

func JwtBearerValidation(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, ProtectedPattern) {
			token := r.Header.Get("Authorization")
			if token == "" {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			} else {
				token := strings.TrimPrefix(token, "Bearer ")
				claims := &jwt.MapClaims{}

				_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
						return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
					}
					return rsaKey, nil
				})

				if err != nil {
					http.Error(w, "Unauthorized", http.StatusUnauthorized)
					return
				}
			}
		}

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

type JwtClaims struct {
	jwt.Claims
}

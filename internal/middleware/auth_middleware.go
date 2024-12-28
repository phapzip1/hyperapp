package middleware

import (
	"context"
	"fmt"
	"hyperapp/misc"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func IsAuthorized(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			misc.ResponseWithError(w, http.StatusUnauthorized, "missing authorization header!")
			return
		}

		tokens := strings.Split(authHeader, " ")
		if len(tokens) != 2 || tokens[0] != "Bearer" {
			misc.ResponseWithError(w, http.StatusUnauthorized, "malformed authorization header!")
		}

		token, err := jwt.Parse(tokens[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return jwt.ParseRSAPublicKeyFromPEM([]byte(os.Getenv("KEYCLOAK_PUBLIC_KEY")))
		})

		if err != nil {
			misc.ResponseWithError(w, http.StatusInternalServerError, "invalid public key!")
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			/// Write user claims to context
			ctx := r.Context()
			ctx = context.WithValue(ctx, "user", claims)

			next(w, r)
		} else {
			fmt.Println(err)
		}

	}
}

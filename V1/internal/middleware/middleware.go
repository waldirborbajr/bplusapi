package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"github.com/waldirborbajr/bplusapi/internal/helper"
	"github.com/waldirborbajr/bplusapi/internal/services"
)

func IsAuthorized(next http.Handler) http.Handler {
	godotenv.Load(".env")
	myKey := []byte(os.Getenv("SECRET_KEY"))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Authorization"] != nil {
			token, err := jwt.Parse(
				r.Header["Authorization"][0],
				func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("there was an error")
					}
					return myKey, nil
				},
			)
			if err != nil {
				payload := services.JsonResponse{
					Error:   true,
					Message: err.Error(),
				}
				_ = helper.WriteJSON(w, http.StatusUnauthorized, payload)
				return
			}
			if token.Valid {
				next.ServeHTTP(w, r)
			}
		} else {
			helper.ErrorJSON(w, errors.New("authorization headers missing"))
		}
	})
}

package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/nyelwa-senguji/ticketing_system_backend/token"
	"github.com/nyelwa-senguji/ticketing_system_backend/utils"
)

const (
	authorizationHeaderKey  = "authorization"
	auhtorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func AuthMiddleware(tokenMaker token.Maker, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header not provided")
			utils.WriteJSON(w, http.StatusUnauthorized, err)
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			utils.WriteJSON(w, http.StatusUnauthorized, err)
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != auhtorizationTypeBearer {
			err := fmt.Errorf("unsupported authorization type %s", authorizationType)
			utils.WriteJSON(w, http.StatusUnauthorized, err)
			return
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			utils.WriteJSON(w, http.StatusUnauthorized, err)
			return
		}

		w.Header().Set(authorizationPayloadKey, payload.Username)

		next.ServeHTTP(w, r)
	})
}

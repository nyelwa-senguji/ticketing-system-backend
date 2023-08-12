package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/nyelwa-senguji/ticketing_system_backend/token"
	"github.com/nyelwa-senguji/ticketing_system_backend/utils"
)

type Middleware interface {
	HeaderMiddleware(next http.Handler) http.Handler
	AuthenticationMiddleware(next http.Handler) http.Handler
}

type middleware struct {
	tokenMaker token.Maker
}

func NewMiddleware() Middleware {
	tokenMaker, _ := token.NewPasetoMaker("f4b49eb23-23ebs567ju-acv78kl2832")
	return &middleware{
		tokenMaker: tokenMaker,
	}
}

const (
	authorizationHeaderKey  = "authorization"
	auhtorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func (m middleware) HeaderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func (m middleware) AuthenticationMiddleware(next http.Handler) http.Handler {
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
		payload, err := m.tokenMaker.VerifyToken(accessToken)
		if err != nil {
			utils.WriteJSON(w, http.StatusUnauthorized, err)
			return
		}

		w.Header().Set(authorizationPayloadKey, payload.Username)

		next.ServeHTTP(w, r)
	})
}

package middleware

import (
	"net/http"
	"strings"
	"workflow-service/transport/util"

	"firebase.google.com/go/v4/auth"
)

// Middleware for authenticating users.
func Authenticate(authClient *auth.Client) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenString := parseBearerToken(r.Header)
			if tokenString == "" {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			ctx := r.Context()
			token, err := authClient.VerifyIDToken(ctx, tokenString)
			if err != nil {
				util.WriteErrResponse(w, http.StatusBadRequest, err)
				return
			}

			ctx = util.WithUserId(ctx, token.UID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// Function for parsing Authorization header in the HTTP request.
func parseBearerToken(h http.Header) string {
	if h == nil {
		return ""
	}
	return strings.TrimPrefix(h.Get("Authorization"), "Bearer ")
}

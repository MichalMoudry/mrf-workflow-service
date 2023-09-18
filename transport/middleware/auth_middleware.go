package middleware

import (
	"net/http"
	"strings"
	"workflow-service/transport/model/ioc"
	"workflow-service/transport/util"
)

// A middleware for handling user's authentication.
func Authenticate(userService ioc.IUserService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token, err := userService.VerifyIdToken(
				r.Context(),
				parseBearerToken(r.Header),
			)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			ctx := util.WithUserId(r.Context(), token)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// Private function for parsing a Bearer token from a HTTP request header.
func parseBearerToken(h http.Header) string {
	if h == nil {
		return ""
	}
	return strings.TrimPrefix(h.Get("Authorization"), "Bearer ")
}

package util

import (
	"context"

	"firebase.google.com/go/v4/auth"
)

// Function for appending an ID of a user to a context.
func WithUserId(ctx context.Context, token *auth.Token) context.Context {
	return context.WithValue(ctx, "user_id", token.Subject)
}

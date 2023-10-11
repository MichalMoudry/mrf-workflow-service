package util

import "context"

type UserId string

// Function
func WithUserId(ctx context.Context, userId string) context.Context {
	var uidKey UserId = "user_id"
	return context.WithValue(ctx, uidKey, userId)
}

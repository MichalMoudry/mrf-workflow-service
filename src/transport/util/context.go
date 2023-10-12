package util

import "context"

type UserId string

const (
	UIDKEY UserId = "user_id"
)

// Function for adding UID to a context.
func WithUserId(ctx context.Context, userId string) context.Context {
	return context.WithValue(ctx, UIDKEY, userId)
}

// Function for retrieving UID from a context.
func GetUserIdFromCtx(ctx context.Context) (string, bool) {
	userId, ok := ctx.Value(UIDKEY).(string)
	return userId, ok
}

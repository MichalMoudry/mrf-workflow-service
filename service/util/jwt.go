package util

import (
	"context"

	"github.com/go-chi/jwtauth/v5"
)

// Function for obtaining a map of claims from a JWT token.
func GetClaimsFromContext(ctx context.Context) (map[string]interface{}, error) {
	_, claims, err := jwtauth.FromContext(ctx)
	if err != nil {
		return map[string]interface{}{}, err
	}
	return claims, nil
}

// Function for obtaining a user id from context.
func GetUserIdFromContext(ctx context.Context) (interface{}, error) {
	claims, err := GetClaimsFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return claims["user_id"], nil
}

package util

import (
	"context"
	"fmt"

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
func GetUserIdFromContext(ctx context.Context) (string, error) {
	claims, err := GetClaimsFromContext(ctx)
	if err != nil {
		return "", err
	}
	return fmt.Sprint(claims["sub"]), nil
}

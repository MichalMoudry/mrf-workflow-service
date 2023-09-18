package service

import (
	"context"

	"firebase.google.com/go/v4/auth"
)

// A structure representing a service providing logic related to users.
type UserService struct {
	AuthClient *auth.Client
}

// Constructor function for the UserService structure.
func NewUserService(authClient *auth.Client) *UserService {
	return &UserService{
		AuthClient: authClient,
	}
}

// Method for validating user's token.
func (srvc UserService) VerifyIdToken(ctx context.Context, token string) (*auth.Token, error) {
	return srvc.AuthClient.VerifyIDToken(ctx, token)
}

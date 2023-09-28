package domain

import "errors"

type ErrorResponse struct {
	Message string `json:"message"`
}

// Error List:
var (
	ErrUserNotAllowed            = errors.New("user not allowed")
	ErrUserNotFound              = errors.New("user not found")
	ErrUnauthorized              = errors.New("unauthorized")
	ErrInvalidPassword           = errors.New("invalid password")
	ErrUserShouldLoginWithGoogle = errors.New("user should login with Google")
	ErrCodeExchangeWrong         = errors.New("code exchange wrong")
	ErrFailedGetGoogleUser       = errors.New("failed to get google user")
	ErrFailedToReadResponse      = errors.New("failed to read response")
	ErrUnexpectedSigningMethod   = errors.New("unexpected signing method")
	ErrInvalidToken              = errors.New("invalid token")
)

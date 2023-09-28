package domain

import (
	"github.com/golang-jwt/jwt/v4"
)

type JwtCustomClaims struct {
	Name     string `json:"name"`
	ID       int    `json:"id"`
	Email    string `json:"email"`
	GoogleId string `json:"google_id"`
	jwt.RegisteredClaims
}

type JwtCustomRefreshClaims struct {
	Name     string `json:"name"`
	ID       int    `json:"id"`
	Email    string `json:"email"`
	GoogleId string `json:"google_id"`
	jwt.RegisteredClaims
}

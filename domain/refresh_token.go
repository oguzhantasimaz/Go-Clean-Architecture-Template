package domain

import (
	"context"

	"github.com/oguzhantasimaz/Go-Clean-Architecture-Template/bootstrap"
)

type RefreshTokenRequest struct {
	RefreshToken string `form:"refreshToken" binding:"required"`
}

type RefreshTokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type RefreshTokenUseCase interface {
	RefreshToken(ctx context.Context, request RefreshTokenRequest, env *bootstrap.Env) (accessToken string, refreshToken string, err error)
}

package usecase

import (
	"context"
	"time"

	"github.com/oguzhantasimaz/Go-Clean-Architecture-Template/bootstrap"
	"github.com/oguzhantasimaz/Go-Clean-Architecture-Template/domain"
	"github.com/oguzhantasimaz/Go-Clean-Architecture-Template/internal/tokenutil"
	"github.com/oguzhantasimaz/Go-Clean-Architecture-Template/repository"

	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type loginUseCase struct {
	userRepository repository.UserRepository
	contextTimeout time.Duration
}

func NewLoginUseCase(userRepository repository.UserRepository, timeout time.Duration) domain.LoginUseCase {
	return &loginUseCase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (lu *loginUseCase) Login(ctx context.Context, request domain.LoginRequest, env *bootstrap.Env) (accessToken string, refreshToken string, err error) {
	var user *domain.User
	user, err = lu.userRepository.GetUserByEmail(ctx, request.Email)
	if err != nil {
		log.Error(err)
		return
	}

	if user.GoogleId != "" {
		log.Error("User should login with Google")
		err = domain.ErrUserShouldLoginWithGoogle
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		log.Error("Invalid password")
		err = domain.ErrInvalidPassword
		return
	}

	accessToken, err = tokenutil.CreateAccessToken(user, env.AccessTokenSecret, env.AccessTokenExpiryHour)
	if err != nil {
		log.Error(err)
		return
	}

	refreshToken, err = tokenutil.CreateRefreshToken(user, env.RefreshTokenSecret, env.RefreshTokenExpiryHour)
	if err != nil {
		log.Error(err)
		return
	}

	return accessToken, refreshToken, nil
}

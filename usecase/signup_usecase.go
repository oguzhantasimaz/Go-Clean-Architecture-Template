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

type signupUseCase struct {
	userRepository repository.UserRepository
	contextTimeout time.Duration
}

func NewSignupUseCase(userRepository repository.UserRepository, timeout time.Duration) domain.SignupUseCase {
	return &signupUseCase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (su *signupUseCase) SignUp(ctx context.Context, request domain.SignupRequest, env *bootstrap.Env) (accessToken string, refreshToken string, err error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		log.Error(err)
		return
	}

	request.Password = string(encryptedPassword)

	user := &domain.User{
		Email:     request.Email,
		Password:  request.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	user, err = su.userRepository.CreateUser(ctx, user)
	if err != nil {
		log.Error(err)
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

	return
}

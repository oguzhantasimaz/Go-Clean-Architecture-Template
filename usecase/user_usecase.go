package usecase

import (
	"context"
	"time"

	"github.com/oguzhantasimaz/Go-Clean-Architecture-Template/domain"
	"github.com/oguzhantasimaz/Go-Clean-Architecture-Template/repository"
)

type userUseCase struct {
	userRepository repository.UserRepository
	contextTimeout time.Duration
}

func NewUserUseCase(userRepository repository.UserRepository, timeout time.Duration) domain.UserUseCase {
	return &userUseCase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (uu *userUseCase) GetUsers(c context.Context) ([]*domain.UserResponse, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	var urs []*domain.UserResponse
	users, err := uu.userRepository.GetUsers(ctx)
	if err != nil {
		return nil, err
	}
	for _, user := range users {
		urs = append(urs, &domain.UserResponse{
			Id:             user.Id,
			GoogleId:       user.GoogleId,
			ProfilePicture: user.ProfilePicture,
			Name:           user.Name,
			Email:          user.Email,
			Phone:          user.Phone,
			CreatedAt:      user.CreatedAt,
		})
	}
	return urs, nil
}

func (uu *userUseCase) GetUserById(c context.Context, id int) (*domain.UserResponse, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	var ur *domain.UserResponse
	user, err := uu.userRepository.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}
	ur = &domain.UserResponse{
		Id:             user.Id,
		GoogleId:       user.GoogleId,
		ProfilePicture: user.ProfilePicture,
		Name:           user.Name,
		Email:          user.Email,
		Phone:          user.Phone,
		CreatedAt:      user.CreatedAt,
	}
	return ur, nil
}

func (uu *userUseCase) UpdateUser(c context.Context, user *domain.User) error {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	return uu.userRepository.UpdateUser(ctx, user)
}

func (uu *userUseCase) DeleteUser(c context.Context, id int) error {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	return uu.userRepository.DeleteUser(ctx, id)
}

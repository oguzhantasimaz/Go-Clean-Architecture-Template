package usecase

import (
	"context"
	"github.com/oguzhantasimaz/Go-Clean-Architecture-Template/bootstrap"
	"github.com/oguzhantasimaz/Go-Clean-Architecture-Template/domain"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type mockUserRepository struct{}

func (m *mockUserRepository) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	return &domain.User{
		Id:        1,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (m *mockUserRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	return &domain.User{
		Id:        1,
		Email:     email,
		Password:  "testPassword",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (m *mockUserRepository) GetUserById(ctx context.Context, id int) (*domain.User, error) {
	return &domain.User{
		Id:        1,
		Email:     "email",
		Password:  "testPassword",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (m *mockUserRepository) GetUsers(ctx context.Context) ([]*domain.User, error) {
	return []*domain.User{
		{
			Id:        1,
			Email:     "email",
			Password:  "testPassword",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}, nil
}

func (m *mockUserRepository) UpdateUser(ctx context.Context, user *domain.User) error {
	return nil
}

func (m *mockUserRepository) DeleteUser(ctx context.Context, id int) error {
	return nil
}

func TestSignUp(t *testing.T) {
	ctx := context.Background()
	timeout := time.Second * 5
	env := &bootstrap.Env{
		AccessTokenSecret:      "testAccessTokenSecret",
		AccessTokenExpiryHour:  24,
		RefreshTokenSecret:     "testRefreshTokenSecret",
		RefreshTokenExpiryHour: 24 * 7,
	}

	// Create a mock user repository
	userRepo := &mockUserRepository{}

	// Create the signupUseCase with the mock user repository
	signupUC := NewSignupUseCase(userRepo, timeout)

	// Test signup request
	request := domain.SignupRequest{
		Email:    "test@example.com",
		Password: "testPassword",
	}

	accessToken, refreshToken, err := signupUC.SignUp(ctx, request, env)

	// Assertions
	assert.NoError(t, err, "Error occurred during signup")
	assert.NotEmpty(t, accessToken, "Access token should not be empty")
	assert.NotEmpty(t, refreshToken, "Refresh token should not be empty")
}

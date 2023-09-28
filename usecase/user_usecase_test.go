package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/oguzhantasimaz/Go-Clean-Architecture-Template/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

var timeNow = time.Now()

// Implement the GetUserById function in the mock
func (m *MockUserRepository) GetUserById(ctx context.Context, id int) (*domain.User, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	} else {
		return args.Get(0).(*domain.User), args.Error(1)
	}
}

func (m *MockUserRepository) GetUsers(ctx context.Context) ([]*domain.User, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	} else {
		return args.Get(0).([]*domain.User), args.Error(1)
	}
}

func (m *MockUserRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	args := m.Called(ctx, email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	} else {
		return args.Get(0).(*domain.User), args.Error(1)
	}
}

func (m *MockUserRepository) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	args := m.Called(ctx, user)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	} else {
		return args.Get(0).(*domain.User), args.Error(1)
	}
}

func (m *MockUserRepository) UpdateUser(ctx context.Context, user *domain.User) error {
	args := m.Called(ctx, user)
	if args.Get(0) == nil {
		return nil
	} else {
		return args.Error(0)
	}
}

func (m *MockUserRepository) DeleteUser(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil
	} else {
		return args.Error(0)
	}
}

func TestGetUserById_Success(t *testing.T) {
	// Create a mock repository
	mockRepo := new(MockUserRepository)

	// Prepare test data
	expectedUser := &domain.User{
		Id:             1,
		GoogleId:       "some-google-id",
		ProfilePicture: "https://example.com/profile-picture.png",
		Name:           "John Doe",
		Email:          "john@example.com",
		Phone:          "1234567890",
		CreatedAt:      timeNow,
	}

	// Set up the mock behavior for the GetUserById function
	mockRepo.On("GetUserById", mock.Anything, 1).Return(expectedUser, nil)

	// Create a userUseCase instance with the mock repository
	uu := NewUserUseCase(mockRepo, 5)

	// Call the GetUserById function
	userResponse, err := uu.GetUserById(context.Background(), 1)
	user := &domain.User{
		Id:             userResponse.Id,
		GoogleId:       userResponse.GoogleId,
		ProfilePicture: userResponse.ProfilePicture,
		Name:           userResponse.Name,
		Email:          userResponse.Email,
		Phone:          userResponse.Phone,
		CreatedAt:      userResponse.CreatedAt,
	}

	// Assert the results
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, expectedUser.Id, user.Id)
	assert.Equal(t, expectedUser.GoogleId, user.GoogleId)
	assert.Equal(t, expectedUser.ProfilePicture, user.ProfilePicture)
	assert.Equal(t, expectedUser.Name, user.Name)
	assert.Equal(t, expectedUser.Email, user.Email)
	assert.Equal(t, expectedUser.Phone, user.Phone)
	assert.Equal(t, expectedUser.CreatedAt, user.CreatedAt)

	// Ensure that the mock repository's GetUserById function was called with the correct arguments
	mockRepo.AssertExpectations(t)
}

func TestGetUserById_Error(t *testing.T) {
	// Create a mock repository
	mockRepo := new(MockUserRepository)

	// Set up the mock behavior for the GetUserById function
	mockRepo.On("GetUserById", mock.Anything, 2).Return(nil, errors.New("user not found"))

	// Create a userUseCase instance with the mock repository
	uu := NewUserUseCase(mockRepo, time.Second*5)

	// Call the GetUserById function
	userResponse, err := uu.GetUserById(context.Background(), 2)

	// Assert the results
	assert.Error(t, err)
	assert.Nil(t, userResponse)

	// Ensure that the mock repository's GetUserById function was called with the correct arguments
	mockRepo.AssertExpectations(t)
}

func TestGetUsers_Success(t *testing.T) {
	// Create a mock repository
	mockRepo := new(MockUserRepository)

	// Prepare test data
	expectedUser := &domain.User{
		Id:             1,
		GoogleId:       "some-google-id",
		ProfilePicture: "https://example.com/profile-picture.png",
		Name:           "John Doe",
		Email:          "john@example.com",
		Phone:          "1234567890",
		CreatedAt:      time.Now(),
	}

	expectedUsers := []*domain.User{expectedUser}

	// Set up the mock behavior for the GetUsers function
	mockRepo.On("GetUsers", mock.Anything).Return(expectedUsers, nil)

	// Create a userUseCase instance with the mock repository
	uu := NewUserUseCase(mockRepo, time.Second*5)

	// Call the GetUsers function
	usersResponse, err := uu.GetUsers(context.Background())
	var users []*domain.User
	for _, user := range usersResponse {
		users = append(users, &domain.User{
			Id:             user.Id,
			GoogleId:       user.GoogleId,
			ProfilePicture: user.ProfilePicture,
			Name:           user.Name,
			Email:          user.Email,
			Phone:          user.Phone,
			CreatedAt:      user.CreatedAt,
		})
	}

	// Assert the results
	assert.NoError(t, err)
	assert.NotNil(t, users)
	assert.Equal(t, expectedUsers, users)

	// Ensure that the mock repository's GetUsers function was called with the correct arguments
	mockRepo.AssertExpectations(t)
}

func TestGetUsers_Error(t *testing.T) {
	// Create a mock repository
	mockRepo := new(MockUserRepository)

	// Set up the mock behavior for the GetUsers function
	mockRepo.On("GetUsers", mock.Anything).Return(nil, errors.New("error getting users"))

	// Create a userUseCase instance with the mock repository
	uu := NewUserUseCase(mockRepo, time.Second*5)

	// Call the GetUsers function
	usersResponse, err := uu.GetUsers(context.Background())

	// Assert the results
	assert.Error(t, err)
	assert.Nil(t, usersResponse)

	// Ensure that the mock repository's GetUsers function was called with the correct arguments
	mockRepo.AssertExpectations(t)
}

func TestUpdateUsers_Success(t *testing.T) {
	// Create a mock repository
	mockRepo := new(MockUserRepository)

	// Prepare test data
	expectedUser := &domain.User{
		Id:             1,
		GoogleId:       "some-google-id",
		ProfilePicture: "https://example.com/profile-picture.png",
		Name:           "John Doe",
		Email:          "john@example.com",
		Phone:          "1234567890",
		CreatedAt:      timeNow,
	}

	// Set up the mock behavior for the UpdateUser function
	mockRepo.On("UpdateUser", mock.Anything, expectedUser).Return(nil)

	// Create a userUseCase instance with the mock repository
	uu := NewUserUseCase(mockRepo, time.Second*5)

	// Call the UpdateUser function
	err := uu.UpdateUser(context.Background(), expectedUser)

	// Assert the results
	assert.NoError(t, err)

	// Ensure that the mock repository's UpdateUser function was called with the correct arguments
	mockRepo.AssertExpectations(t)
}

func TestUpdateUsers_Error(t *testing.T) {
	// Create a mock repository
	mockRepo := new(MockUserRepository)

	// Prepare test data
	expectedUser := &domain.User{
		Id:             1,
		GoogleId:       "some-google-id",
		ProfilePicture: "https://example.com/profile-picture.png",
		Name:           "John Doe",
		Email:          "john@example.com",
		Phone:          "1234567890",
		CreatedAt:      timeNow,
	}

	// Set up the mock behavior for the UpdateUser function
	mockRepo.On("UpdateUser", mock.Anything, expectedUser).Return(errors.New("error updating user"))

	// Create a userUseCase instance with the mock repository
	uu := NewUserUseCase(mockRepo, time.Second*5)

	// Call the UpdateUser function
	err := uu.UpdateUser(context.Background(), expectedUser)

	// Assert the results
	assert.Error(t, err)

	// Ensure that the mock repository's UpdateUser function was called with the correct arguments
	mockRepo.AssertExpectations(t)
}

func TestDeleteUsers_Success(t *testing.T) {
	// Create a mock repository
	mockRepo := new(MockUserRepository)

	// Prepare test data
	expectedUser := &domain.User{
		Id:             1,
		GoogleId:       "some-google-id",
		ProfilePicture: "https://example.com/profile-picture.png",
		Name:           "John Doe",
		Email:          "john@example.com",
		Phone:          "1234567890",
		CreatedAt:      timeNow,
	}

	// Set up the mock behavior for the DeleteUser function
	mockRepo.On("DeleteUser", mock.Anything, expectedUser.Id).Return(nil)

	// Create a userUseCase instance with the mock repository
	uu := NewUserUseCase(mockRepo, time.Second*5)

	// Call the DeleteUser function
	err := uu.DeleteUser(context.Background(), expectedUser.Id)

	// Assert the results
	assert.NoError(t, err)

	// Ensure that the mock repository's DeleteUser function was called with the correct arguments
	mockRepo.AssertExpectations(t)
}

func TestDeleteUsers_Error(t *testing.T) {
	// Create a mock repository
	mockRepo := new(MockUserRepository)

	// Prepare test data
	expectedUser := &domain.User{
		Id:             1,
		GoogleId:       "some-google-id",
		ProfilePicture: "https://example.com/profile-picture.png",
		Name:           "John Doe",
		Email:          "john@example.com",
		Phone:          "1234567890",
		CreatedAt:      timeNow,
	}

	// Set up the mock behavior for the DeleteUser function
	mockRepo.On("DeleteUser", mock.Anything, expectedUser.Id).Return(errors.New("error deleting user"))

	// Create a userUseCase instance with the mock repository
	uu := NewUserUseCase(mockRepo, time.Second*5)

	// Call the DeleteUser function
	err := uu.DeleteUser(context.Background(), expectedUser.Id)

	// Assert the results
	assert.Error(t, err)

	// Ensure that the mock repository's DeleteUser function was called with the correct arguments
	mockRepo.AssertExpectations(t)
}

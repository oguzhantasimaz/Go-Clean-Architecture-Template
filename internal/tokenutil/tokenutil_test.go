package tokenutil_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/oguzhantasimaz/Go-Clean-Architecture-Template/domain"
	"github.com/oguzhantasimaz/Go-Clean-Architecture-Template/internal/tokenutil"
)

func TestCreateAccessToken(t *testing.T) {
	// Test user information
	user := &domain.User{
		Name:     "John Doe",
		GoogleId: "google123",
		Email:    "john@example.com",
		Id:       123,
	}

	secret := "testAccessTokenSecret"
	expiry := 1 // 1 hour

	// Create an access token
	accessToken, err := tokenutil.CreateAccessToken(user, secret, expiry)

	// Assertions
	assert.NoError(t, err, "Error occurred while creating access token")
	assert.NotEmpty(t, accessToken, "Access token should not be empty")
}

func TestCreateRefreshToken(t *testing.T) {
	// Test user information
	user := &domain.User{
		Name:     "John Doe",
		GoogleId: "google123",
		Email:    "john@example.com",
		Id:       123,
	}

	secret := "testRefreshTokenSecret"
	expiry := 24 * 7 // 1 week

	// Create a refresh token
	refreshToken, err := tokenutil.CreateRefreshToken(user, secret, expiry)

	// Assertions
	assert.NoError(t, err, "Error occurred while creating refresh token")
	assert.NotEmpty(t, refreshToken, "Refresh token should not be empty")
}

func TestIsAuthorized(t *testing.T) {
	secret := "testAccessTokenSecret"
	expiry := 1 // 1 hour

	// Create a test token
	user := &domain.User{
		Name:     "John Doe",
		GoogleId: "google123",
		Email:    "john@example.com",
		Id:       123,
	}
	accessToken, _ := tokenutil.CreateAccessToken(user, secret, expiry)

	// Check if token is authorized
	authorized, err := tokenutil.IsAuthorized(accessToken, secret)

	// Assertions
	assert.NoError(t, err, "Error occurred while checking authorization")
	assert.True(t, authorized, "Token should be authorized")
}

func TestExtractIDFromToken(t *testing.T) {
	secret := "testAccessTokenSecret"
	expiry := 1 // 1 hour

	// Create a test token
	user := &domain.User{
		Name:     "John Doe",
		GoogleId: "google123",
		Email:    "john@example.com",
		Id:       123,
	}
	accessToken, _ := tokenutil.CreateAccessToken(user, secret, expiry)

	// Extract ID from the token
	id, err := tokenutil.ExtractIDFromToken(accessToken, secret)

	// Assertions
	assert.NoError(t, err, "Error occurred while extracting ID from token")
	assert.Equal(t, user.Id, id, "Extracted ID should match user's ID")
}

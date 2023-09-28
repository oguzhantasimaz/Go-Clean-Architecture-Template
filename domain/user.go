package domain

import (
	"context"
	"time"
)

type User struct {
	Id             int       `json:"id" db:"id"`
	GoogleId       string    `json:"google_id" db:"google_id"`
	ProfilePicture string    `json:"profile_picture" db:"profile_picture"`
	Name           string    `json:"name" db:"name"`
	Password       string    `json:"password" db:"password"`
	Email          string    `json:"email" db:"email"`
	Phone          string    `json:"phone" db:"phone"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}

type UserResponse struct {
	Id             int       `json:"id" db:"id"`
	GoogleId       string    `json:"google_id" db:"google_id"`
	ProfilePicture string    `json:"profile_picture" db:"profile_picture"`
	Name           string    `json:"name" db:"name"`
	Email          string    `json:"email" db:"email"`
	Phone          string    `json:"phone" db:"phone"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
}

type UserUseCase interface {
	GetUserById(c context.Context, id int) (*UserResponse, error)
	GetUsers(c context.Context) ([]*UserResponse, error)
	UpdateUser(c context.Context, user *User) error
	DeleteUser(c context.Context, id int) error
}

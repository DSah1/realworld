package user

import (
	"awesomeProject/internal/model"
	"awesomeProject/internal/request"
	"awesomeProject/internal/response"
)

type Store interface {
	GetByID(uint) (*model.User, error)
	GetByEmail(string) (*model.User, error)
	GetByUsername(string2 string) (*model.User, error)
	Create(user *model.User) error
	Update(user *model.User) error

	IsFollower(uint, uint) (bool, error)
	AddFollower(*model.User, uint) error
	RemoveFollower(*model.User, uint) error
	GetFollows(userID uint) ([]model.User, error)
}

type Service interface {
	Registration(*request.UserRegisterRequest) (*response.UserResponse, error)
	Login(*request.UserLoginRequest) (*response.UserResponse, error)
	CurrentUser(uint) (*response.UserResponse, error)
	UpdateUser(uint, *request.UserUpdateRequest) (*response.UserResponse, error)
}

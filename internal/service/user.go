package service

import (
	"awesomeProject/internal/model"
	"awesomeProject/internal/request"
	"awesomeProject/internal/response"
	"awesomeProject/internal/user"
	"github.com/gofiber/fiber/v2"
)

type UserService struct {
	userStore user.Store
}

func NewUserService(us user.Store) *UserService {
	return &UserService{userStore: us}
}

func (us *UserService) Registration(r *request.UserRegisterRequest) (*response.UserResponse, error) {
	var u model.User

	u.Username = r.User.Username
	u.Email = r.User.Email

	if err := u.HashPassword(r.User.Password); err != nil {
		return nil, err
	}

	if err := us.userStore.Create(&u); err != nil {
		return nil, err
	}

	return response.NewUserResponse(&u), nil
}

func (us *UserService) Login(r *request.UserLoginRequest) (*response.UserResponse, error) {
	foundUser, err := us.userStore.GetByEmail(r.User.Email)

	if err != nil {
		return nil, err
	}

	if !foundUser.CheckPassword(r.User.Password) {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Bad email or password")
	}

	return response.NewUserResponse(foundUser), nil
}

func (us *UserService) CurrentUser(userId uint) (*response.UserResponse, error) {
	u, err := us.userStore.GetByID(userId)
	if err != nil {
		return nil, err
	}
	return response.NewUserResponse(u), nil
}

func (us *UserService) UpdateUser(userId uint, r *request.UserUpdateRequest) (*response.UserResponse, error) {
	u, err := us.userStore.GetByID(userId)
	if err != nil {
		return nil, err
	}
	if err := r.Bind(u); err != nil {
		return nil, err
	}

	if err := us.userStore.Update(u); err != nil {
		return nil, err
	}

	return response.NewUserResponse(u), nil

}

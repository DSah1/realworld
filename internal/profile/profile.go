package profile

import (
	"awesomeProject/internal/response"
)

type Service interface {
	GetProfile(username string, userId uint) (*response.ProfileResponse, error)
	Follow(username string, userId uint) (*response.ProfileResponse, error)
	Unfollow(username string, userId uint) (*response.ProfileResponse, error)
}

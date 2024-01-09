package service

import (
	"awesomeProject/internal/response"
	"awesomeProject/internal/user"
)

type ProfileService struct {
	userStore user.Store
}

func NewProfileService(us user.Store) *ProfileService {
	return &ProfileService{userStore: us}
}

func (ps *ProfileService) GetProfile(username string, userId uint) (*response.ProfileResponse, error) {

	u, err := ps.userStore.GetByUsername(username)
	if err != nil {
		return nil, err
	}
	isFollower, err := ps.userStore.IsFollower(u.ID, userId)
	if err != nil {
		return nil, err
	}
	return response.NewProfileResponse(u, isFollower), nil
}

func (ps *ProfileService) Follow(username string, userId uint) (*response.ProfileResponse, error) {
	u, err := ps.userStore.GetByUsername(username)
	if err != nil {
		return nil, err
	}

	if err := ps.userStore.AddFollower(u, userId); err != nil {
		return nil, err
	}
	isFollower, err := ps.userStore.IsFollower(u.ID, userId)
	if err != nil {
		return nil, err
	}
	return response.NewProfileResponse(u, isFollower), nil
}

func (ps *ProfileService) Unfollow(username string, userId uint) (*response.ProfileResponse, error) {
	u, err := ps.userStore.GetByUsername(username)
	if err != nil {
		return nil, err
	}

	if err := ps.userStore.RemoveFollower(u, userId); err != nil {
		return nil, err
	}
	isFollower, err := ps.userStore.IsFollower(u.ID, userId)
	if err != nil {
		return nil, err
	}
	return response.NewProfileResponse(u, isFollower), nil
}

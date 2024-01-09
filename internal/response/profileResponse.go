package response

import (
	"awesomeProject/internal/model"
)

type ProfileResponse struct {
	*Profile `json:"profile"`
}

type Profile struct {
	Username  string  `json:"username"`
	Bio       *string `json:"bio"`
	Image     *string `json:"image"`
	Following bool    `json:"following"`
}

func NewProfileResponse(u *model.User, isFollower bool) *ProfileResponse {
	r := new(ProfileResponse)
	r.Profile = NewProfile(u, isFollower)
	return r
}

func NewProfile(u *model.User, isFollower bool) *Profile {
	p := new(Profile)
	p.Following = isFollower
	p.Bio = u.Bio
	p.Username = u.Username
	p.Image = u.Image

	return p
}

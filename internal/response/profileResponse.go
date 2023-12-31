package response

import (
	"awesomeProject/internal/model"
	"awesomeProject/internal/user"
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

func NewProfileResponse(u *model.User, us user.Store, userID uint) *ProfileResponse {
	r := new(ProfileResponse)
	r.Profile = NewProfile(u, us, userID)
	return r
}

func NewProfile(u *model.User, us user.Store, userID uint) *Profile {
	p := new(Profile)
	p.Following, _ = us.IsFollower(u.ID, userID)
	p.Bio = u.Bio
	p.Username = u.Username
	p.Image = u.Image

	return p
}

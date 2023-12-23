package response

import (
	"awesomeProject/model"
	"awesomeProject/user"
)

type ProfileResponse struct {
	Profile struct {
		Username  string  `json:"username"`
		Bio       *string `json:"bio"`
		Image     *string `json:"image"`
		Following bool    `json:"following"`
	} `json:"profile"`
}

func NewProfileResponse(u *model.User, us user.Store, userID uint) *ProfileResponse {
	r := new(ProfileResponse)

	r.Profile.Following, _ = us.IsFollower(u.ID, userID)
	r.Profile.Bio = u.Bio
	r.Profile.Username = u.Username
	r.Profile.Image = u.Image
	return r
}

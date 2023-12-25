package response

import (
	"awesomeProject/internal/model"
	"awesomeProject/utils"
)

type UserResponse struct {
	User struct {
		Email    string  `json:"email"`
		Token    string  `json:"token"`
		Username string  `json:"username"`
		Bio      *string `json:"bio"`
		Image    *string `json:"image"`
	} `json:"user"`
}

func NewUserResponse(u *model.User) *UserResponse {
	r := new(UserResponse)
	r.User.Username = u.Username
	r.User.Email = u.Email
	r.User.Image = u.Image
	r.User.Bio = u.Bio
	r.User.Token = utils.GenerateJWT(u.ID)

	return r
}

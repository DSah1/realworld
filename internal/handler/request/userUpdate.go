package request

import (
	"awesomeProject/internal/model"
	"github.com/gofiber/fiber/v2"
)

type UserUpdateRequest struct {
	User struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Username string `json:"username"`
		Bio      string `json:"bio"`
		Image    string `json:"image"`
	} `json:"user"`
}

func (r *UserUpdateRequest) Populate(u *model.User) {
	r.User.Username = u.Username
	r.User.Email = u.Email
	r.User.Password = u.Password
	if u.Bio != nil {
		r.User.Bio = *u.Bio
	}
	if u.Image != nil {
		r.User.Image = *u.Image
	}
}

func (r *UserUpdateRequest) Bind(c *fiber.Ctx, u *model.User) error {
	if err := c.BodyParser(r); err != nil {
		return err
	}
	u.Username = r.User.Username
	u.Email = r.User.Email
	if r.User.Password != u.Password {
		h, err := u.HashPassword(r.User.Password)
		if err != nil {
			return err
		}
		u.Password = h
	}
	u.Bio = &r.User.Bio
	u.Image = &r.User.Image
	return nil
}

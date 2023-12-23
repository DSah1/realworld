package request

import (
	"awesomeProject/model"
	"github.com/gofiber/fiber/v2"
)

type UserRegisterRequest struct {
	User struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	} `json:"user"`
}

func (r *UserRegisterRequest) Bind(c *fiber.Ctx, u *model.User) error {
	if err := c.BodyParser(r); err != nil {
		return err
	}

	u.Username = r.User.Username
	u.Email = r.User.Email
	h, err := u.HashPassword(r.User.Password)
	if err != nil {
		return err
	}
	u.Password = h
	return nil
}

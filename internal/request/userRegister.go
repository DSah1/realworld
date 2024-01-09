package request

import (
	"github.com/gofiber/fiber/v2"
)

type UserRegisterRequest struct {
	User struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	} `json:"user"`
}

func (r *UserRegisterRequest) Bind(c *fiber.Ctx) error {
	if err := c.BodyParser(r); err != nil {
		return err
	}
	return nil
}

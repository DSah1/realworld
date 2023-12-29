package request

import "github.com/gofiber/fiber/v2"

type UserLoginRequest struct {
	User struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	} `json:"user"`
}

func (r *UserLoginRequest) Bind(c *fiber.Ctx) error {
	if err := c.BodyParser(r); err != nil {
		return err
	}

	return nil
}

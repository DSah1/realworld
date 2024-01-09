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

func (r *UserUpdateRequest) ParseUser(c *fiber.Ctx) error {
	if err := c.BodyParser(r); err != nil {
		return err
	}
	return nil
}

func (r *UserUpdateRequest) Bind(u *model.User) error {
	updateIfNotEmpty(&u.Username, r.User.Username)
	updateIfNotEmpty(&u.Email, r.User.Email)

	if r.User.Bio != "" {
		u.Bio = &r.User.Bio
	}

	if r.User.Image != "" {
		u.Image = &r.User.Image
	}

	if r.User.Password != "" {
		if err := u.HashPassword(r.User.Password); err != nil {
			return err
		}
	}

	return nil
}

func updateIfNotEmpty(current *string, newValue string) {
	if newValue != "" {
		*current = newValue
	}
}

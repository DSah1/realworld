package handler

import (
	"awesomeProject/internal/model"
	"awesomeProject/internal/request"
	"awesomeProject/internal/response"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
)

func (h *Handler) Register(c *fiber.Ctx) error {

	var user model.User

	req := request.UserRegisterRequest{}

	if err := req.Bind(c, &user); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusUnprocessableEntity).SendString(err.Error())
	}

	err := user.HashPassword(req.User.Password)
	if err != nil {
		return err
	}

	if err := h.userStore.Create(&user); err != nil {
		return c.Status(http.StatusUnprocessableEntity).SendString(err.Error())
	}

	return c.Status(http.StatusCreated).JSON(response.NewUserResponse(&user))
}

func (h *Handler) Login(c *fiber.Ctx) error {

	req := request.UserLoginRequest{}

	err := req.Bind(c)

	if err != nil {
		return err
	}

	foundUser, err := h.userStore.GetByEmail(req.User.Email)

	if err != nil {
		return err
	}

	if !foundUser.CheckPassword(req.User.Password) {
		return c.Status(http.StatusUnauthorized).SendString("Bad email or password")
	}
	return c.Status(http.StatusAccepted).JSON(response.NewUserResponse(foundUser))
}

func (h *Handler) CurrentUser(c *fiber.Ctx) error {
	user, err := h.userStore.GetByID(getUserIDByToken(c))
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(response.NewUserResponse(user))
}

func (h *Handler) UpdateUser(c *fiber.Ctx) error {

	user, err := h.userStore.GetByID(getUserIDByToken(c))

	if err != nil {
		return err
	}
	req := request.UserUpdateRequest{}

	req.Populate(user)

	if err := req.Bind(c, user); err != nil {
		return c.SendStatus(http.StatusUnprocessableEntity)
	}
	if req.User.Password != user.Password {
		err := user.HashPassword(req.User.Password)
		if err != nil {
			return err
		}
	}

	return c.Status(http.StatusOK).JSON(response.NewUserResponse(user))
}

func getUserIDByToken(c *fiber.Ctx) uint {

	var user *jwt.Token
	l := c.Locals("user")
	if l == nil {
		return 0
	}

	user = l.(*jwt.Token)
	id := uint((user.Claims.(jwt.MapClaims)["id"]).(float64))

	return id
}

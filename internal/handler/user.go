package handler

import (
	"awesomeProject/internal/request"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
)

func (h *Handler) Register(c *fiber.Ctx) error {

	req := request.UserRegisterRequest{}

	if err := req.Bind(c); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusUnprocessableEntity).SendString(err.Error())
	}
	userResponse, err := h.userService.Registration(&req)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	return c.Status(http.StatusCreated).JSON(userResponse)
}

func (h *Handler) Login(c *fiber.Ctx) error {

	req := request.UserLoginRequest{}

	err := req.Bind(c)

	if err != nil {
		return err
	}

	userResponse, err := h.userService.Login(&req)

	if err != nil {
		return err
	}

	return c.Status(http.StatusAccepted).JSON(userResponse)
}

func (h *Handler) CurrentUser(c *fiber.Ctx) error {
	userResponse, err := h.userService.CurrentUser(getUserIDByToken(c))
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(userResponse)
}

func (h *Handler) UpdateUser(c *fiber.Ctx) error {

	req := request.UserUpdateRequest{}
	if err := req.ParseUser(c); err != nil {
		return c.SendStatus(http.StatusUnprocessableEntity)
	}
	userResponse, err := h.userService.UpdateUser(getUserIDByToken(c), &req)

	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(userResponse)
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

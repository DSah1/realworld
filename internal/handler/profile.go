package handler

import (
	"awesomeProject/internal/response"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func (h *Handler) GetProfile(c *fiber.Ctx) error {

	user, err := h.userStore.GetByUsername(c.Params("username"))
	if err != nil {
		return err
	}

	currentUserID := getUserIDByToken(c)

	return c.Status(http.StatusOK).JSON(response.NewProfileResponse(user, h.userStore, currentUserID))
}

func (h *Handler) Follow(c *fiber.Ctx) error {
	user, err := h.userStore.GetByUsername(c.Params("username"))
	if err != nil {
		return err
	}

	currentUserID := getUserIDByToken(c)

	err = h.userStore.AddFollower(user, currentUserID)
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(response.NewProfileResponse(user, h.userStore, currentUserID))
}

func (h *Handler) Unfollow(c *fiber.Ctx) error {
	user, err := h.userStore.GetByUsername(c.Params("username"))
	if err != nil {
		return err
	}

	currentUserID := getUserIDByToken(c)

	err = h.userStore.RemoveFollower(user, currentUserID)
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(response.NewProfileResponse(user, h.userStore, currentUserID))
}

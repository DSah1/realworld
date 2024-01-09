package handler

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func (h *Handler) GetProfile(c *fiber.Ctx) error {
	currentUserID := getUserIDByToken(c)
	profileResponse, err := h.profileService.GetProfile(c.Params("username"), currentUserID)

	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(profileResponse)
}

func (h *Handler) Follow(c *fiber.Ctx) error {
	currentUserID := getUserIDByToken(c)
	profileResponse, err := h.profileService.Follow(c.Params("username"), currentUserID)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(profileResponse)
}

func (h *Handler) Unfollow(c *fiber.Ctx) error {
	currentUserID := getUserIDByToken(c)
	profileResponse, err := h.profileService.Unfollow(c.Params("username"), currentUserID)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(profileResponse)
}

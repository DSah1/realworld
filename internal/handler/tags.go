package handler

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func (h *Handler) Tags(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(h.articleService.GetTags())
}

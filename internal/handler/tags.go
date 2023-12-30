package handler

import (
	"awesomeProject/internal/response"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func (h *Handler) Tags(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(response.NewTagsResponse(h.articleStore.GetTags()))
}

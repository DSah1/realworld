package handler

import (
	"awesomeProject/internal/request"
	"awesomeProject/utils"
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func (h *Handler) ListArticles(c *fiber.Ctx) error {
	userID := getUserIDByToken(c)
	query := c.Queries()
	tag := query["tag"]
	//TODO: NOT IMPLEMENTED
	//author := query["author"]
	//favorited := query["favorited"]

	limit := utils.IntFromQuery(query, "limit", 20)
	offset := utils.IntFromQuery(query, "offset", 0)

	resp, err := h.articleService.ListArticles(tag, "", "", limit, offset, userID)
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(resp)
}

func (h *Handler) Feed(c *fiber.Ctx) error {
	userID := getUserIDByToken(c)
	query := c.Queries()

	limit := utils.IntFromQuery(query, "limit", 20)
	offset := utils.IntFromQuery(query, "offset", 0)

	resp, err := h.articleService.GetFeed(limit, offset, userID)

	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(resp)
}

func (h *Handler) CreateArticle(c *fiber.Ctx) error {
	userID := getUserIDByToken(c)
	req := new(request.CreateArticleRequest)

	if err := req.ParseArticle(c); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return c.Status(http.StatusBadRequest).SendString("Duplicate Key")
		}
		return err
	}

	resp, err := h.articleService.CreateArticle(req, userID)
	if err != nil {
		return err
	}

	return c.Status(http.StatusCreated).JSON(resp)
}

func (h *Handler) GetArticle(c *fiber.Ctx) error {
	slug := c.Params("slug")
	userID := getUserIDByToken(c)

	resp, err := h.articleService.GetArticle(slug, userID)

	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(resp)
}

func (h *Handler) UpdateArticle(c *fiber.Ctx) error {
	slug := c.Params("slug")
	userID := getUserIDByToken(c)

	req := request.UpdateArticleRequest{}

	if err := req.ParseArticle(c); err != nil {
		return err
	}

	resp, err := h.articleService.UpdateArticle(slug, userID, &req)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(resp)
}

func (h *Handler) DeleteArticle(c *fiber.Ctx) error {
	slug := c.Params("slug")
	userID := getUserIDByToken(c)

	resp, err := h.articleService.DeleteArticle(slug, userID)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(resp)
}

func (h *Handler) FavoriteArticle(c *fiber.Ctx) error {
	userID := getUserIDByToken(c)
	slug := c.Params("slug")

	resp, err := h.articleService.FavoriteArticle(slug, userID)

	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(resp)
}

func (h *Handler) UnfavoriteArticle(c *fiber.Ctx) error {
	userID := getUserIDByToken(c)
	slug := c.Params("slug")

	resp, err := h.articleService.UnfavoriteArticle(slug, userID)
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(resp)
}

func (h *Handler) Comment(c *fiber.Ctx) error {
	userID := getUserIDByToken(c)
	slug := c.Params("slug")

	req := request.CreateCommentRequest{}

	if err := req.ParseBody(c); err != nil {
		return err
	}
	resp, err := h.articleService.CommentArticle(slug, userID, &req)
	if err != nil {
		return err
	}

	return c.Status(http.StatusCreated).JSON(resp)
}

func (h *Handler) DeleteComment(c *fiber.Ctx) error {
	userID := getUserIDByToken(c)
	slug := c.Params("slug")
	commentId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}
	if err := h.articleService.DeleteComment(slug, userID, uint(commentId)); err != nil {
		return err
	}

	return c.SendStatus(http.StatusOK)
}

func (h *Handler) AllComments(c *fiber.Ctx) error {
	userID := getUserIDByToken(c)
	slug := c.Params("slug")

	resp, err := h.articleService.AllComments(slug, userID)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(resp)
}

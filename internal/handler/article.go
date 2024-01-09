package handler

import (
	"awesomeProject/internal/model"
	"awesomeProject/internal/request"
	"awesomeProject/internal/response"
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

	article, err := h.articleStore.GetBySlug(slug)

	if err != nil {
		return err
	}

	if userID != article.AuthorID {
		return c.SendStatus(http.StatusUnauthorized)
	}

	if err := h.articleStore.Delete(article); err != nil {
		return err
	}

	isFollower, err := h.userStore.IsFollower(article.AuthorID, userID)
	if err != nil {
		return err
	}

	inFavorites := h.articleStore.IsUserInFavorites(article.ID, userID)

	return c.Status(http.StatusOK).JSON(response.NewArticleResponse(article, isFollower, inFavorites))
}

func (h *Handler) FavoriteArticle(c *fiber.Ctx) error {
	userID := getUserIDByToken(c)
	article, err := h.articleStore.GetBySlug(c.Params("slug"))

	if err != nil {
		return err
	}
	isFollower, err := h.userStore.IsFollower(article.AuthorID, userID)
	if err != nil {
		return err
	}

	if h.articleStore.IsUserInFavorites(article.ID, userID) {
		return c.Status(http.StatusOK).JSON(response.NewArticleResponse(article, isFollower, true))
	}

	if err := h.articleStore.AddFavorite(article, userID); err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(response.NewArticleResponse(article, isFollower, true))
}

func (h *Handler) UnfavoriteArticle(c *fiber.Ctx) error {
	userID := getUserIDByToken(c)
	article, err := h.articleStore.GetBySlug(c.Params("slug"))

	if err != nil {
		return err
	}
	isFollower, err := h.userStore.IsFollower(article.AuthorID, userID)
	if err != nil {
		return err
	}

	if !h.articleStore.IsUserInFavorites(article.ID, userID) {
		return c.Status(http.StatusOK).JSON(response.NewArticleResponse(article, isFollower, false))
	}

	if err := h.articleStore.RemoveFavorite(article, userID); err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(response.NewArticleResponse(article, isFollower, false))
}

func (h *Handler) Comment(c *fiber.Ctx) error {
	userID := getUserIDByToken(c)
	article, err := h.articleStore.GetBySlug(c.Params("slug"))

	if err != nil {
		return err
	}

	req := request.CreateCommentRequest{}
	comment := model.Comment{}

	if err := req.Bind(c, &comment, userID, article); err != nil {
		return err
	}

	if err := h.articleStore.CreateComment(&comment); err != nil {
		return err
	}

	isFollow, err := h.userStore.IsFollower(comment.UserID, userID)
	if err != nil {
		return err
	}

	return c.Status(http.StatusCreated).JSON(response.NewCommentResponse(&comment, isFollow))
}

func (h *Handler) DeleteComment(c *fiber.Ctx) error {
	userID := getUserIDByToken(c)
	article, err := h.articleStore.GetBySlug(c.Params("slug"))

	if err != nil {
		return err
	}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	comment, err := h.articleStore.GetCommentByID(uint(id))

	if comment.UserID != userID {
		return c.Status(http.StatusForbidden).SendString("You cannot delete comments that were not made by you")
	}

	if comment.ArticleID != article.ID {
		return c.Status(http.StatusForbidden).SendString("This comment is not related to slug")
	}

	if err != nil {
		return err
	}

	if err := h.articleStore.DeleteComment(comment); err != nil {
		return err
	}
	return c.SendStatus(http.StatusOK)
}

func (h *Handler) AllComments(c *fiber.Ctx) error {
	userID := getUserIDByToken(c)
	comments, err := h.articleStore.GetCommentsForArticle(c.Params("slug"))

	if err != nil {
		return err
	}
	isFollows := make([]bool, len(comments))
	for i := range comments {
		isFollow, err := h.userStore.IsFollower(comments[i].UserID, userID)
		if err != nil {
			return err
		}
		isFollows[i] = isFollow
	}

	return c.Status(http.StatusOK).JSON(response.NewMultipleCommentResponse(comments, isFollows))
}

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

	if tag != "" {
		articles, err := h.articleStore.ListByTag(limit, offset, tag)
		if err != nil {
			return err
		}
		articleResponse := response.NewMultiArticleResponse(articles, h.articleStore, h.userStore, userID)
		//TODO: NOT IMPLEMENTED
		return c.Status(http.StatusOK).JSON(articleResponse)
	}

	articles, err := h.articleStore.List(limit, offset)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(response.NewMultiArticleResponse(articles, h.articleStore, h.userStore, userID))
}

func (h *Handler) CreateArticle(c *fiber.Ctx) error {
	userID := getUserIDByToken(c)
	user, err := h.userStore.GetByID(userID)

	var article model.Article

	req := new(request.CreateArticleRequest)
	err = req.Bind(c, &article, user)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return c.Status(http.StatusBadRequest).SendString("Duplicate Key")
		}
		return err
	}

	err = h.articleStore.Create(&article)
	if err != nil {
		return err
	}

	return c.Status(http.StatusCreated).JSON(response.NewArticleResponse(&article, h.articleStore, h.userStore, userID))
}

func (h *Handler) GetArticle(c *fiber.Ctx) error {
	slug := c.Params("slug")
	userID := getUserIDByToken(c)

	article, err := h.articleStore.GetBySlug(slug)

	if err != nil {
		return err
	}
	if article == nil {
		return c.Status(http.StatusNotFound).SendString("Article not found")
	}

	res := response.NewArticleResponse(article, h.articleStore, h.userStore, userID)

	return c.Status(http.StatusOK).JSON(res)
}

func (h *Handler) UpdateArticle(c *fiber.Ctx) error {
	slug := c.Params("slug")
	userID := getUserIDByToken(c)

	article, err := h.articleStore.GetBySlug(slug)

	if article == nil {
		return c.Status(http.StatusNotFound).SendString("Article not found")
	}

	if err != nil {
		return err
	}

	if userID != article.AuthorID {
		return c.Status(http.StatusUnauthorized).SendString("You are not authorized to update this article, not author.")
	}

	r := request.UpdateArticleRequest{}
	r.Populate(article)

	if err := r.Bind(c, article); err != nil {
		return err
	}

	if err := h.articleStore.Update(article); err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(response.NewArticleResponse(article, h.articleStore, h.userStore, userID))
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

	return c.Status(http.StatusOK).JSON(response.NewArticleResponse(article, h.articleStore, h.userStore, userID))
}

func (h *Handler) FavoriteArticle(c *fiber.Ctx) error {
	userID := getUserIDByToken(c)
	article, err := h.articleStore.GetBySlug(c.Params("slug"))

	if err != nil {
		return err
	}

	if h.articleStore.IsUserInFavorites(article.ID, userID) {
		return c.Status(http.StatusOK).JSON(response.NewArticleResponse(article, h.articleStore, h.userStore, userID))
	}

	if err := h.articleStore.AddFavorite(article, userID); err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(response.NewArticleResponse(article, h.articleStore, h.userStore, userID))
}

func (h *Handler) UnfavoriteArticle(c *fiber.Ctx) error {
	userID := getUserIDByToken(c)
	article, err := h.articleStore.GetBySlug(c.Params("slug"))

	if err != nil {
		return err
	}

	if !h.articleStore.IsUserInFavorites(article.ID, userID) {
		return c.Status(http.StatusOK).JSON(response.NewArticleResponse(article, h.articleStore, h.userStore, userID))
	}

	if err := h.articleStore.RemoveFavorite(article, userID); err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(response.NewArticleResponse(article, h.articleStore, h.userStore, userID))
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

	return c.Status(http.StatusCreated).JSON(response.NewCommentResponse(&comment, h.userStore, userID))
}

func (h *Handler) DeleteComment(c *fiber.Ctx) error {
	//userID := getUserIDByToken(c)
	//article, err := h.articleStore.GetBySlug(c.Params("slug"))

	//if err != nil {
	//	return err
	//}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	comment, err := h.articleStore.GetCommentByID(uint(id))
	if err != nil {
		return err
	}

	if err := h.articleStore.DeleteComment(comment); err != nil {
		return err
	}
	return c.SendStatus(http.StatusOK)
}

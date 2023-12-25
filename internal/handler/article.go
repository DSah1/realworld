package handler

import (
	"awesomeProject/internal/model"
	"awesomeProject/internal/request"
	"awesomeProject/internal/response"
	"awesomeProject/utils"
	"github.com/gofiber/fiber/v2"
	"net/http"
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

	res := response.NewArticleResponse(article, h.articleStore, h.userStore, userID)

	return c.Status(http.StatusOK).JSON(res)
}

func (h *Handler) UpdateArticle(c *fiber.Ctx) error {
	slug := c.Params("slug")
	userID := getUserIDByToken(c)

	article, err := h.articleStore.GetBySlug(slug)

	if err != nil {
		return err
	}

	if userID != article.AuthorID {
		return c.SendStatus(http.StatusUnauthorized)
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

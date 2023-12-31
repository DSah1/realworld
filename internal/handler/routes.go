package handler

import (
	"awesomeProject/utils"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

func (h *Handler) RegisterRoutes(r *fiber.App) {
	v1 := r.Group("/api")

	jwtmidware := jwtware.New(jwtware.Config{
		SigningKey: utils.SecretString,
		AuthScheme: "Token",
	})

	guestUsers := v1.Group("/users")
	guestUsers.Post("", h.Register)
	guestUsers.Post("/login", h.Login)

	user := v1.Group("/user", jwtmidware)
	user.Get("", h.CurrentUser)
	user.Put("", h.UpdateUser)

	profile := v1.Group("/profiles", jwtmidware)
	profile.Get("/:username", h.GetProfile)
	profile.Post("/:username/follow", h.Follow)
	profile.Post("/:username/unfollow", h.Unfollow)

	articles := v1.Group("/articles", jwtmidware)
	articles.Get("", h.ListArticles)
	articles.Get("/feed", h.Feed)

	articles.Post("", h.CreateArticle)
	articles.Get("/:slug", h.GetArticle)
	articles.Put("/:slug", h.UpdateArticle)
	articles.Delete("/:slug", h.DeleteArticle)

	articles.Post("/:slug/favorite", h.FavoriteArticle)
	articles.Delete("/:slug/favorite", h.UnfavoriteArticle)

	articles.Post("/:slug/comments", h.Comment)
	articles.Delete("/:slug/comments/:id", h.DeleteComment)
	articles.Get("/:slug/comments", h.AllComments)

	tags := v1.Group("tags")
	tags.Get("", h.Tags)
}

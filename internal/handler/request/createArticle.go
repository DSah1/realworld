package request

import (
	"awesomeProject/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gosimple/slug"
)

type CreateArticleRequest struct {
	Article struct {
		Title       string   `json:"title"`
		Description string   `json:"description"`
		Body        string   `json:"body"`
		TagList     []string `json:"tagList,omitempty"`
	} `json:"article"`
}

func (r *CreateArticleRequest) Bind(c *fiber.Ctx, a *model.Article, u *model.User) error {
	if err := c.BodyParser(r); err != nil {
		return err
	}

	a.Title = r.Article.Title
	a.Slug = slug.Make(r.Article.Title)
	a.Description = r.Article.Description
	a.Author = *u
	a.AuthorID = u.ID

	if r.Article.TagList != nil {
		for _, tag := range r.Article.TagList {
			a.Tags = append(a.Tags, model.Tag{Tag: tag})
		}
	}

	return nil
}

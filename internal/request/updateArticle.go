package request

import (
	"awesomeProject/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gosimple/slug"
)

type UpdateArticleRequest struct {
	Article struct {
		Title       string   `json:"title"`
		Description string   `json:"description"`
		Body        string   `json:"body"`
		TagList     []string `json:"tagList,omitempty"`
	} `json:"article"`
}

func (u *UpdateArticleRequest) Populate(a *model.Article) {
	u.Article.Title = a.Title
	u.Article.Description = a.Description
	u.Article.Body = a.Body
	u.Article.TagList = a.ExtractTags()
}

func (u *UpdateArticleRequest) ParseArticle(c *fiber.Ctx) error {
	if err := c.BodyParser(u); err != nil {
		return err
	}
	return nil
}

func (u *UpdateArticleRequest) Bind(a *model.Article) error {
	if u.Article.Title != "" {
		a.Title = u.Article.Title
		a.Slug = slug.Make(u.Article.Title)
	}

	updateIfNotEmpty(&a.Title, u.Article.Title)
	updateIfNotEmpty(&a.Body, u.Article.Body)

	if u.Article.TagList != nil {
		a.Tags = make([]model.Tag, len(u.Article.TagList))
		for i := range u.Article.TagList {
			a.Tags[i] = model.Tag{Tag: u.Article.TagList[i]}
		}
	}
	return nil
}

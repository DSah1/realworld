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

func (u *UpdateArticleRequest) Bind(c *fiber.Ctx, a *model.Article) error {
	if err := c.BodyParser(u); err != nil {
		return err
	}

	a.Title = u.Article.Title
	a.Slug = slug.Make(u.Article.Title)
	a.Description = u.Article.Description
	a.Body = u.Article.Body
	a.Tags = []model.Tag{}
	if u.Article.TagList != nil {
		for _, tag := range u.Article.TagList {
			a.Tags = append(a.Tags, model.Tag{Tag: tag})
		}
	}

	return nil
}

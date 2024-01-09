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

func (r *CreateArticleRequest) ParseArticle(c *fiber.Ctx) error {
	if err := c.BodyParser(r); err != nil {
		return err
	}
	return nil
}

func (r *CreateArticleRequest) Bind(a *model.Article, uid uint) {
	a.Title = r.Article.Title
	a.Slug = slug.Make(r.Article.Title)
	a.Description = r.Article.Description
	a.AuthorID = uid

	if r.Article.TagList != nil {
		for _, tag := range r.Article.TagList {
			a.Tags = append(a.Tags, model.Tag{Tag: tag})
		}
	}

}

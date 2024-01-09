package request

import (
	"awesomeProject/internal/model"
	"github.com/gofiber/fiber/v2"
)

type CreateCommentRequest struct {
	Comment struct {
		Body string `json:"body"`
	} `json:"comment"`
}

func (r *CreateCommentRequest) ParseBody(c *fiber.Ctx) error {
	if err := c.BodyParser(r); err != nil {
		return err
	}
	return nil
}

func (r *CreateCommentRequest) Bind(comment *model.Comment, userID uint, article *model.Article) error {

	comment.Body = r.Comment.Body
	comment.UserID = userID
	comment.Article = *article
	comment.ArticleID = article.ID

	return nil
}

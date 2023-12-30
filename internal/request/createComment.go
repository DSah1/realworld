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

func (r *CreateCommentRequest) Bind(c *fiber.Ctx, comment *model.Comment, userID uint, article *model.Article) error {
	if err := c.BodyParser(r); err != nil {
		return err
	}

	comment.Body = r.Comment.Body
	comment.UserID = userID
	comment.Article = *article
	comment.ArticleID = article.ID

	return nil
}

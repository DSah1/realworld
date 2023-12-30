package response

import (
	"awesomeProject/internal/model"
	"awesomeProject/internal/user"
	"awesomeProject/utils"
)

type SingleComment struct {
	Comment `json:"comment"`
}

type MultipleComments struct {
	Comments []Comment `json:"comments"`
}

type Comment struct {
	Id        uint            `json:"id"`
	CreatedAt string          `json:"createdAt"`
	UpdatedAt string          `json:"updatedAt"`
	Body      string          `json:"body"`
	Author    ProfileResponse `json:"author"`
}

func NewCommentResponse(comment *model.Comment, us user.Store, userID uint) *SingleComment {
	resComment := new(SingleComment)
	resComment.Comment = *assignToComment(comment, us, userID)

	return resComment
}

func NewMultipleCommentResponse(comments []model.Comment, us user.Store, userID uint) *MultipleComments {
	resComment := new(MultipleComments)
	for i := range comments {
		resComment.Comments = append(resComment.Comments, *assignToComment(&comments[i], us, userID))
	}

	return resComment
}

func assignToComment(cmt *model.Comment, us user.Store, userID uint) *Comment {
	resComment := new(Comment)

	resComment.Id = cmt.ID
	resComment.Body = cmt.Body
	resComment.Author = *NewProfileResponse(&cmt.User, us, userID)
	resComment.CreatedAt = cmt.CreatedAt.Format(utils.ISO8601)
	resComment.UpdatedAt = cmt.UpdatedAt.Format(utils.ISO8601)

	return resComment
}

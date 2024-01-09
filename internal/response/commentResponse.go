package response

import (
	"awesomeProject/internal/model"
	"awesomeProject/utils"
)

type SingleComment struct {
	Comment `json:"comment"`
}

type MultipleComments struct {
	Comments []Comment `json:"comments"`
}

type Comment struct {
	Id        uint     `json:"id"`
	CreatedAt string   `json:"createdAt"`
	UpdatedAt string   `json:"updatedAt"`
	Body      string   `json:"body"`
	Author    *Profile `json:"author"`
}

func NewCommentResponse(comment *model.Comment, isFollower bool) *SingleComment {
	resComment := new(SingleComment)
	resComment.Comment = *assignToComment(comment, isFollower)

	return resComment
}

func NewMultipleCommentResponse(comments []model.Comment, isFollowers []bool) *MultipleComments {
	resComment := new(MultipleComments)
	for i := range comments {
		resComment.Comments = append(resComment.Comments, *assignToComment(&comments[i], isFollowers[i]))
	}

	return resComment
}

func assignToComment(cmt *model.Comment, isFollower bool) *Comment {
	resComment := new(Comment)

	resComment.Id = cmt.ID
	resComment.Body = cmt.Body
	resComment.Author = NewProfile(&cmt.User, isFollower)
	resComment.CreatedAt = cmt.CreatedAt.Format(utils.ISO8601)
	resComment.UpdatedAt = cmt.UpdatedAt.Format(utils.ISO8601)

	return resComment
}

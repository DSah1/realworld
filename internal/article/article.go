package article

import (
	"awesomeProject/internal/model"
	"awesomeProject/internal/request"
	"awesomeProject/internal/response"
)

type Store interface {
	List(limit, offset int) ([]model.Article, error)
	ListByTag(limit, offset int, tag string) ([]model.Article, error)
	ListByAuthor(limit, offset int, author string) ([]model.Article, error)
	Feed(int, int, uint) ([]model.Article, error)
	IsUserInFavorites(articleID uint, userID uint) bool
	Create(article *model.Article) error
	Update(article *model.Article) error
	Delete(article *model.Article) error
	GetBySlug(slug string) (*model.Article, error)
	AddFavorite(a *model.Article, userID uint) error
	RemoveFavorite(a *model.Article, userID uint) error
	GetTags() []model.Tag
	CreateComment(*model.Comment) error
	DeleteComment(*model.Comment) error
	GetCommentByID(uint) (*model.Comment, error)
	GetCommentsForArticle(string) ([]model.Comment, error)
}

type Service interface {
	ListArticles(tag, author, favorited string, limit, offset int, userId uint) (*response.MultipleArticle, error)
	GetFeed(limit, offset int, userId uint) (*response.MultipleArticle, error)
	CreateArticle(request *request.CreateArticleRequest, userId uint) (*response.SingleArticle, error)
	GetArticle(slug string, userId uint) (*response.SingleArticle, error)
	UpdateArticle(slug string, userId uint, request *request.UpdateArticleRequest) (*response.SingleArticle, error)
	DeleteArticle(slug string, userId uint) (*response.SingleArticle, error)
	FavoriteArticle(slug string, userId uint) (*response.SingleArticle, error)
	UnfavoriteArticle(slug string, userId uint) (*response.SingleArticle, error)

	CommentArticle(slug string, userId uint, createReq *request.CreateCommentRequest) (*response.SingleComment, error)
	DeleteComment(slug string, userId, commentId uint) error
	AllComments(slug string, userId uint) (*response.MultipleComments, error)
}

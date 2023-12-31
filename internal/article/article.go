package article

import "awesomeProject/internal/model"

type Store interface {
	List(limit, offset int) ([]model.Article, error)
	ListByTag(limit, offset int, tag string) ([]model.Article, error)
	ListByAuthor(limit, offset int, author model.User) ([]model.Article, error)
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

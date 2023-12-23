package article

import "awesomeProject/model"

type Store interface {
	List(limit, offset int) ([]model.Article, error)
	ListByTag(limit, offset int, tag string) ([]model.Article, error)
	ListByAuthor(limit, offset int, author model.User) ([]model.Article, error)
	IsUserInFavorites(articleID uint, userID uint) (bool, error)
	Create(article *model.Article) error
	Update(article *model.Article) error
	Delete(article *model.Article) error
	GetBySlug(slug string) (*model.Article, error)
}

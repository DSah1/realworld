package store

import (
	"awesomeProject/internal/model"
	"errors"
	"gorm.io/gorm"
)

type ArticleStore struct {
	db *gorm.DB
}

func NewArticleStore(db *gorm.DB) *ArticleStore {
	return &ArticleStore{db: db}
}

func (as *ArticleStore) List(limit, offset int) ([]model.Article, error) {
	var articles []model.Article
	err := as.db.Limit(limit).Offset(offset).Find(&articles).Error
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (as *ArticleStore) ListByTag(limit, offset int, tag string) ([]model.Article, error) {
	var articles []model.Article

	err := as.db.Limit(limit).Offset(offset).Where("tag = ?", tag).Find(&articles).Error
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (as *ArticleStore) ListByAuthor(limit, offset int, author model.User) ([]model.Article, error) {
	var articles []model.Article

	err := as.db.Limit(limit).Offset(offset).Where("user = ?", author).Find(&articles).Error
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (as *ArticleStore) IsUserInFavorites(articleID uint, userID uint) (bool, error) {
	var count int64
	err := as.db.Model(&model.Article{}).
		Joins("JOIN favorites ON favorites.article.id = articles.id").
		Where("articles.id = ?", articleID).
		Where("favorites.user.id", userID).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (as *ArticleStore) Create(a *model.Article) error {
	return as.db.Create(a).Error
}

func (as *ArticleStore) Update(a *model.Article) error {
	return as.db.Model(a).Updates(a).Error
}

func (as *ArticleStore) Delete(a *model.Article) error {
	return as.db.Delete(&a).Error
}

func (as *ArticleStore) GetBySlug(slug string) (*model.Article, error) {
	var article model.Article

	if err := as.db.Where("slug = ?", slug).First(&article).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &article, nil
}

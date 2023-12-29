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
	err := as.db.
		Preload("Tags").
		Preload("Favorites").
		Limit(limit).
		Offset(offset).
		Order("created_at desc").
		Find(&articles).
		Error

	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (as *ArticleStore) ListByTag(limit, offset int, tag string) ([]model.Article, error) {
	var articles []model.Article

	err := as.db.
		Preload("Tags", "tag = ?", tag).
		Preload("Favorites").
		Limit(limit).
		Offset(offset).
		Order("created_at desc").
		Find(&articles).Error

	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (as *ArticleStore) ListByAuthor(limit, offset int, author model.User) ([]model.Article, error) {
	var articles []model.Article

	err := as.db.
		Preload("Tags").
		Preload("Favorites").
		Where("author = ?", author).
		Limit(limit).
		Offset(offset).
		Order("created_at desc").
		Find(&articles).Error
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (as *ArticleStore) IsUserInFavorites(articleID uint, userID uint) bool {
	var count int64

	err := as.db.Model(&model.Article{}).
		Joins("JOIN favorites ON favorites.article_id = articles.id").
		Where("articles.id = ? AND favorites.user_id = ?", articleID, userID).
		Count(&count).Error

	if err != nil {
		return false
	}

	return count > 0
}

func (as *ArticleStore) Create(a *model.Article) error {
	return as.db.Create(a).Error
}

func (as *ArticleStore) Update(a *model.Article) error {
	return as.db.Model(a).Updates(a).Error
}

func (as *ArticleStore) Delete(a *model.Article) error {
	if err := as.db.Model(&a).Association("Tags").Clear(); err != nil {
		return err
	}

	if err := as.db.Model(&a).Association("Favorites").Clear(); err != nil {
		return err
	}

	return as.db.Unscoped().Delete(&a).Error
}

func (as *ArticleStore) GetBySlug(slug string) (*model.Article, error) {
	var article model.Article

	if err := as.db.Where("slug = ?", slug).First(&article).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}

	return &article, nil
}

func (as *ArticleStore) AddFavorite(a *model.Article, userID uint) error {
	usr := model.User{}
	usr.ID = userID
	return as.db.Model(a).Association("Favorites").Append(&usr)
}

func (as *ArticleStore) RemoveFavorite(a *model.Article, userID uint) error {
	usr := model.User{}
	usr.ID = userID
	return as.db.Model(a).Association("Favorites").Delete(&usr)
}

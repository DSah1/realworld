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
		Preload("Comments").
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
		Preload("Comments").
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

func (as *ArticleStore) Feed(limit, offset int, userID uint) ([]model.Article, error) {
	var feed []model.Article

	if err := as.db.Model(&model.Article{}).Preload("Tags").Preload("Favorites").
		Joins("JOIN users on users.id = articles.author_id").
		Joins("JOIN follows on follows.following_id = users.id").
		Where("follows.follower_id = ?", userID).Order("articles.created_at DESC").
		Limit(limit).Offset(offset).
		Find(&feed).Error; err != nil {
		return nil, err
	}
	if len(feed) == 0 {
		feed = make([]model.Article, 0)
	}

	return feed, nil
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

	if err := as.db.
		Where(&model.Article{Slug: slug}).Preload("Favorites").
		Preload("Tags", func(db *gorm.DB) *gorm.DB {
			return db.Order("tag asc")
		}).
		Preload("Author").First(&article).Error; err != nil {
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

func (as *ArticleStore) GetTags() []model.Tag {
	var tags []model.Tag
	err := as.db.Model(&model.Tag{}).Find(&tags).Error

	if err != nil || tags == nil {
		return make([]model.Tag, 0)
	}
	return tags
}

func (as *ArticleStore) CreateComment(comment *model.Comment) error {
	return as.db.Create(comment).Error
}

func (as *ArticleStore) DeleteComment(comment *model.Comment) error {
	return as.db.Unscoped().Delete(&comment).Error
}

func (as *ArticleStore) GetCommentByID(commentID uint) (*model.Comment, error) {
	var comment model.Comment
	err := as.db.Model(&comment).Where("id = ?", commentID).Find(&comment).Error
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

func (as *ArticleStore) GetCommentsForArticle(s string) ([]model.Comment, error) {
	var comments []model.Comment

	err := as.db.Model(comments).Preload("User").Where(model.Comment{Article: model.Article{Slug: s}}).Find(&comments).Error
	if err != nil {
		return make([]model.Comment, 0), err
	}
	return comments, nil
}

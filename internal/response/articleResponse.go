package response

import (
	"awesomeProject/article"
	"awesomeProject/internal/model"
	"awesomeProject/internal/user"
	"awesomeProject/utils"
)

type Article struct {
	Slug           string           `json:"slug"`
	Title          string           `json:"title"`
	Description    string           `json:"description"`
	Body           string           `json:"body"`
	TagList        []string         `json:"tagList"`
	CreatedAt      string           `json:"createdAt"`
	UpdatedAt      string           `json:"updatedAt"`
	Favorited      bool             `json:"favorited"`
	FavoritesCount int              `json:"favoritesCount"`
	Author         *ProfileResponse `json:"author"`
}

type SingleArticle struct {
	Article *Article `json:"article"`
}

type MultipleArticle struct {
	Articles      []Article `json:"articles"`
	ArticlesCount int       `json:"articlesCount"`
}

func NewMultiArticleResponse(articles []model.Article, as article.Store, us user.Store, userID uint) *MultipleArticle {
	resArticle := new(MultipleArticle)

	for _, a := range articles {
		resArticle.Articles = append(resArticle.Articles, *assignToArticle(a, as, us, userID))
		resArticle.ArticlesCount++
	}

	return resArticle
}

func NewArticleResponse(article *model.Article, as article.Store, us user.Store, userID uint) *SingleArticle {
	resArticle := new(SingleArticle)

	resArticle.Article = assignToArticle(*article, as, us, userID)

	return resArticle
}

func assignToArticle(article model.Article, as article.Store, us user.Store, userID uint) *Article {
	resArticle := new(Article)

	author, err := us.GetByID(article.AuthorID)
	if err != nil {
		return nil
	}
	article.Author = *author

	resArticle.Slug = article.Slug
	resArticle.Title = article.Title
	resArticle.Description = article.Description
	resArticle.Body = article.Body
	resArticle.TagList = article.ExtractTags()
	resArticle.CreatedAt = article.CreatedAt.Format(utils.ISO8601)
	resArticle.UpdatedAt = article.UpdatedAt.Format(utils.ISO8601)
	resArticle.Favorited, _ = as.IsUserInFavorites(article.ID, userID)
	resArticle.FavoritesCount = len(article.Favorites)
	resArticle.Author = NewProfileResponse(&article.Author, us, userID)

	return resArticle
}

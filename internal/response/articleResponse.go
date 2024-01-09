package response

import (
	"awesomeProject/internal/model"
	"awesomeProject/utils"
)

type Article struct {
	Slug           string   `json:"slug"`
	Title          string   `json:"title"`
	Description    string   `json:"description"`
	Body           string   `json:"body"`
	TagList        []string `json:"tagList"`
	CreatedAt      string   `json:"createdAt"`
	UpdatedAt      string   `json:"updatedAt"`
	Favorited      bool     `json:"favorited"`
	FavoritesCount int      `json:"favoritesCount"`
	Author         *Profile `json:"author"`
}

type SingleArticle struct {
	Article *Article `json:"article"`
}

type MultipleArticle struct {
	Articles      []Article `json:"articles"`
	ArticlesCount int       `json:"articlesCount"`
}

func NewMultiArticleResponse(articles []model.Article, isFollowers, inFavorites []bool) *MultipleArticle {
	resArticle := new(MultipleArticle)
	resArticle.Articles = make([]Article, len(articles))

	for i, a := range articles {
		resArticle.Articles[i] = *assignToArticle(&a, isFollowers[i], inFavorites[i])
		resArticle.ArticlesCount++
	}

	return resArticle
}

func NewArticleResponse(article *model.Article, isFollower, inFavorite bool) *SingleArticle {
	resArticle := new(SingleArticle)

	resArticle.Article = assignToArticle(article, isFollower, inFavorite)

	return resArticle
}

func assignToArticle(article *model.Article, isFollower, inFavorite bool) *Article {
	resArticle := new(Article)

	resArticle.Slug = article.Slug
	resArticle.Title = article.Title
	resArticle.Description = article.Description
	resArticle.Body = article.Body
	resArticle.TagList = article.ExtractTags()
	resArticle.CreatedAt = article.CreatedAt.Format(utils.ISO8601)
	resArticle.UpdatedAt = article.UpdatedAt.Format(utils.ISO8601)
	resArticle.Favorited = inFavorite
	resArticle.FavoritesCount = len(article.Favorites)
	resArticle.Author = NewProfile(&article.Author, isFollower)

	return resArticle
}

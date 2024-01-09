package service

import (
	"awesomeProject/internal/article"
	"awesomeProject/internal/model"
	"awesomeProject/internal/request"
	"awesomeProject/internal/response"
	"awesomeProject/internal/user"
	"errors"
)

type ArticleService struct {
	userStore    user.Store
	articleStore article.Store
}

func NewArticleService(us user.Store, as article.Store) *ArticleService {
	return &ArticleService{userStore: us, articleStore: as}
}

func (as *ArticleService) ListArticles(tag, author, favorited string, limit, offset int, userId uint) (*response.MultipleArticle, error) {
	var articles []model.Article
	if tag != "" {
		var err error
		articles, err = as.articleStore.ListByTag(limit, offset, tag)
		if err != nil {
			return nil, err
		}

	} else if author != "" {
		var err error
		articles, err = as.articleStore.ListByAuthor(limit, offset, author)
		if err != nil {
			return nil, err
		}

	} else if favorited != "" {
		panic("NOT IMPLEMENTED! WTF???")
	} else {
		var err error
		articles, err = as.articleStore.List(limit, offset)
		if err != nil {
			return nil, err
		}
	}
	isFollowers, inFavorites := as.getBooleans(articles, userId)

	return response.NewMultiArticleResponse(articles, isFollowers, inFavorites), nil
}

func (as *ArticleService) GetFeed(limit, offset int, userId uint) (*response.MultipleArticle, error) {
	feed, err := as.articleStore.Feed(limit, offset, userId)
	if err != nil {
		return nil, err
	}

	isFollowers, inFavorites := as.getBooleans(feed, userId)

	return response.NewMultiArticleResponse(feed, isFollowers, inFavorites), nil
}

func (as *ArticleService) CreateArticle(r *request.CreateArticleRequest, userId uint) (*response.SingleArticle, error) {
	var a model.Article
	r.Bind(&a, userId)

	if err := as.articleStore.Create(&a); err != nil {
		return nil, err
	}

	author, err := as.userStore.GetByID(userId)
	if err != nil {
		return nil, err
	}

	a.Author = *author

	return response.NewArticleResponse(&a, false, false), nil
}

func (as *ArticleService) GetArticle(slug string, userId uint) (*response.SingleArticle, error) {
	a, err := as.articleStore.GetBySlug(slug)

	if err != nil {
		return nil, err
	}

	isFollower, err := as.userStore.IsFollower(a.AuthorID, userId)
	if err != nil {
		return nil, err
	}

	inFavorites := as.articleStore.IsUserInFavorites(a.ID, userId)

	return response.NewArticleResponse(a, isFollower, inFavorites), nil
}

func (as *ArticleService) UpdateArticle(slug string, userId uint, req *request.UpdateArticleRequest) (*response.SingleArticle, error) {
	a, err := as.articleStore.GetBySlug(slug)

	if err != nil {
		return nil, err
	}

	if userId != a.AuthorID {
		return nil, errors.New("you are not authorized to update this article, not author")
	}

	if err := req.Bind(a); err != nil {
		return nil, err
	}

	if err := as.articleStore.Update(a); err != nil {
		return nil, err
	}

	inFavorites := as.articleStore.IsUserInFavorites(a.ID, userId)

	return response.NewArticleResponse(a, false, inFavorites), nil
}

func (as *ArticleService) DeleteArticle(slug string, userId uint) (*response.SingleArticle, error) {
	//TODO implement me
	panic("implement me")
}

func (as *ArticleService) FavoriteArticle(slug string, userId uint) (*response.SingleArticle, error) {
	//TODO implement me
	panic("implement me")
}

func (as *ArticleService) UnfavoriteArticle(slug string, userId uint) (*response.SingleArticle, error) {
	//TODO implement me
	panic("implement me")
}

func (as *ArticleService) CommentArticle(slug string, userId uint) (*response.SingleComment, error) {
	//TODO implement me
	panic("implement me")
}

func (as *ArticleService) DeleteComment(slug string, userId uint) (*response.SingleComment, error) {
	//TODO implement me
	panic("implement me")
}

func (as *ArticleService) AllComments(slug string, userId uint) (*response.MultipleComments, error) {
	//TODO implement me
	panic("implement me")
}

func (as *ArticleService) getBooleans(articles []model.Article, userId uint) (isFollowers, inFavorites []bool) {
	isFollowers = make([]bool, len(articles))
	inFavorites = make([]bool, len(articles))

	for i := range articles {
		isFollower, _ := as.userStore.IsFollower(articles[i].AuthorID, userId)
		inFavorite := as.articleStore.IsUserInFavorites(articles[i].ID, userId)

		isFollowers[i] = isFollower
		inFavorites[i] = inFavorite
	}
	return
}

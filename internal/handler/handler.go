package handler

import (
	"awesomeProject/internal/article"
	"awesomeProject/internal/profile"
	"awesomeProject/internal/user"
)

type Handler struct {
	userService    user.Service
	profileService profile.Service
	articleService article.Service

	//DEPRECATED
	userStore user.Store
	//DEPRECATED
	articleStore article.Store
}

func NewHandler(us user.Service, ps profile.Service, as article.Service, uStore user.Store, aStore article.Store) *Handler {
	return &Handler{
		userService:    us,
		profileService: ps,
		articleService: as,
		userStore:      uStore,
		articleStore:   aStore,
	}
}

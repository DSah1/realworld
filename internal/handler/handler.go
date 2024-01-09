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
}

func NewHandler(us user.Service, ps profile.Service, as article.Service) *Handler {
	return &Handler{
		userService:    us,
		profileService: ps,
		articleService: as,
	}
}

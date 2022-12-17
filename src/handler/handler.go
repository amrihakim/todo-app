package handler

import (
	"todo-app/src/repositories"
	"todo-app/src/usecase"

	"gorm.io/gorm"
)

type Handler struct {
	useCase usecase.UseCase
}

func InitializeHandler(db *gorm.DB) *Handler {
	return &Handler{
		useCase: usecase.InitializeUseCase(repositories.InitialRepositories(db)),
	}
}

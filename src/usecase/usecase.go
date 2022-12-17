package usecase

import (
	"context"
	"todo-app/src/models"
	"todo-app/src/repositories"
)

type usecase struct {
	repositories repositories.Repositories
}

type UseCase interface {
	// Activity
	GetActivityList(ctx context.Context) ([]*models.Activity, error)
	GetActivity(ctx context.Context, id string) (*models.Activity, error)
	CreateActivity(ctx context.Context, body *models.Activity) (*models.Activity, error)
	DeleteActivity(ctx context.Context, id string) error
	UpdateActivity(ctx context.Context, body *models.Activity, id string) (*models.Activity, error)

	// Todo
	GetTodoList(ctx context.Context, param string) ([]*models.Todo, error)
	GetTodo(ctx context.Context, id string) (*models.Todo, error)
	CreateTodo(ctx context.Context, body *models.Todo) (*models.Todo, error)
	DeleteTodo(ctx context.Context, id string) error
	UpdateTodo(ctx context.Context, body *models.Todo, id string) (*models.Todo, error)
}

func InitializeUseCase(repo repositories.Repositories) UseCase {
	return &usecase{
		repositories: repo,
	}
}

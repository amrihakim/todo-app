package repositories

import (
	"context"
	"todo-app/src/models"

	"gorm.io/gorm"
)

type repositories struct {
	db *gorm.DB
}

type Repositories interface {
	GetActivityList(ctx context.Context) (resp []*models.Activity, err error)
	GetActivity(ctx context.Context, id string) (resp *models.Activity, err error)
	CreateActivity(ctx context.Context, body *models.Activity) (resp *models.Activity, err error)
	DeleteActivity(ctx context.Context, id string) (err error)
	UpdateActivity(ctx context.Context, body *models.Activity, id string) (resp *models.Activity, err error)

	GetTodoList(ctx context.Context, param string) (resp []*models.Todo, err error)
	GetTodo(ctx context.Context, id string) (resp *models.Todo, err error)
	CreateTodo(ctx context.Context, body *models.Todo) (resp *models.Todo, err error)
	DeleteTodo(ctx context.Context, id string) (err error)
	UpdateTodo(ctx context.Context, body *models.Todo, id string) (resp *models.Todo, err error)
}

func InitialRepositories(db *gorm.DB) Repositories {
	return &repositories{
		db: db,
	}
}

package usecase

import (
	"context"
	"errors"
	"todo-app/src/models"
)

func (uc *usecase) GetTodoList(ctx context.Context, param string) ([]*models.Todo, error) {
	data, err := uc.repositories.GetTodoList(ctx, param)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (uc *usecase) GetTodo(ctx context.Context, id string) (*models.Todo, error) {
	data, err := uc.repositories.GetTodo(ctx, id)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (uc *usecase) CreateTodo(ctx context.Context, body *models.Todo) (*models.Todo, error) {
	if body.Title == "" {
		return nil, errors.New("title cannot be null")
	}

	if body.ActivityGroupID == "" {
		return nil, errors.New("activity_group_id cannot be null")
	}

	data, err := uc.repositories.CreateTodo(ctx, body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (uc *usecase) DeleteTodo(ctx context.Context, id string) error {
	err := uc.repositories.DeleteTodo(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (uc *usecase) UpdateTodo(ctx context.Context, body *models.Todo, id string) (*models.Todo, error) {

	data, err := uc.repositories.UpdateTodo(ctx, body, id)
	if err != nil {
		return nil, err
	}

	return data, nil
}

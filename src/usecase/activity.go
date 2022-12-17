package usecase

import (
	"context"
	"errors"
	"todo-app/src/models"
)

func (uc *usecase) GetActivityList(ctx context.Context) ([]*models.Activity, error) {
	data, err := uc.repositories.GetActivityList(ctx)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (uc *usecase) GetActivity(ctx context.Context, id string) (*models.Activity, error) {
	data, err := uc.repositories.GetActivity(ctx, id)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (uc *usecase) CreateActivity(ctx context.Context, body *models.Activity) (*models.Activity, error) {
	if body.Title == "" {
		return nil, errors.New("title cannot be null")
	}

	data, err := uc.repositories.CreateActivity(ctx, body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (uc *usecase) DeleteActivity(ctx context.Context, id string) error {
	err := uc.repositories.DeleteActivity(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (uc *usecase) UpdateActivity(ctx context.Context, body *models.Activity, id string) (*models.Activity, error) {
	if body.Title == "" {
		return nil, errors.New("title cannot be null")
	}

	data, err := uc.repositories.UpdateActivity(ctx, body, id)
	if err != nil {
		return nil, err
	}

	return data, nil
}

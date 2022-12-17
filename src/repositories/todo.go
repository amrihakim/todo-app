package repositories

import (
	"context"
	"errors"
	"fmt"
	"todo-app/src/models"

	"gorm.io/gorm"
)

func (r *repositories) GetTodoList(ctx context.Context, param string) (resp []*models.Todo, err error) {
	var model []*models.Todo
	db := r.db.Model(&model)

	if param != "" {
		db.Where("activity_group_id", param)
	}

	db = db.Order("created_at DESC")
	if err = db.Find(&model).Error; err != nil {
		return nil, err
	}

	return model, nil
}

func (r *repositories) GetTodo(ctx context.Context, id string) (resp *models.Todo, err error) {
	var model *models.Todo
	db := r.db.Model(&model)

	if id != "" {
		db.Where("id", id)
	}

	if err = db.First(&model).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = fmt.Errorf("Todo with ID %v Not Found", id)
			return nil, err
		}
		return nil, err
	}

	return model, nil
}

func (r *repositories) CreateTodo(ctx context.Context, body *models.Todo) (resp *models.Todo, err error) {
	if err := r.db.Model(models.Todo{}).Create(&body).Scan(&resp).Error; err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *repositories) DeleteTodo(ctx context.Context, id string) (err error) {
	var model *models.Todo
	db := r.db.Model(&model)

	if err = db.Where("id", id).First(&model).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = fmt.Errorf("Todo with ID %v Not Found", id)
			return err
		}
		return err
	}

	db.Delete(&model)
	return nil
}

func (r *repositories) UpdateTodo(ctx context.Context, body *models.Todo, id string) (resp *models.Todo, err error) {
	var model *models.Todo
	db := r.db.Model(&model)

	if err = db.Where("id", id).First(&model).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = fmt.Errorf("Todo with ID %v Not Found", id)
			return nil, err
		}
		return nil, err
	}

	db.Updates(body)
	return model, nil
}

package repositories

import (
	"context"
	"errors"
	"fmt"
	"todo-app/src/models"

	"gorm.io/gorm"
)

func (r *repositories) GetActivityList(ctx context.Context) (resp []*models.Activity, err error) {
	var model []*models.Activity
	db := r.db.Model(&model)

	db = db.Order("created_at DESC")
	if err = db.Find(&model).Error; err != nil {
		return nil, err
	}

	return model, nil
}

func (r *repositories) GetActivity(ctx context.Context, id string) (resp *models.Activity, err error) {
	var model *models.Activity
	db := r.db.Model(&model)

	if id != "" {
		db.Where("id", id)
	}

	if err = db.First(&model).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = fmt.Errorf("Activity with ID %v Not Found", id)
			return nil, err
		}
		return nil, err
	}

	return model, nil
}

func (r *repositories) CreateActivity(ctx context.Context, body *models.Activity) (resp *models.Activity, err error) {
	if err := r.db.Model(models.Activity{}).Create(&body).Scan(&resp).Error; err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *repositories) DeleteActivity(ctx context.Context, id string) (err error) {
	var model *models.Activity
	db := r.db.Model(&model)

	if err = db.Where("id", id).First(&model).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = fmt.Errorf("Activity with ID %v Not Found", id)
			return err
		}
		return err
	}

	db.Delete(&model)
	return nil
}

func (r *repositories) UpdateActivity(ctx context.Context, body *models.Activity, id string) (resp *models.Activity, err error) {
	var model *models.Activity
	db := r.db.Model(&model)

	if err = db.Where("id", id).First(&model).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = fmt.Errorf("Activity with ID %v Not Found", id)
			return nil, err
		}
		return nil, err
	}

	db.Updates(body)
	return model, nil
}

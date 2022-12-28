package models

import (
	"time"

	"gorm.io/gorm"
)

type Activity struct {
	CreatedAt time.Time      `gorm:"type:timestamp;column:created_at;default:current_timestamp" json:"created_at"`
	UpdatedAt *time.Time     `gorm:"type:timestamp;column:updated_at" json:"updated_at"`
	ID        uint64         `gorm:"autoIncrement;primaryKey;column:id" json:"id"`
	Title     string         `gorm:"type:varchar(255);column:title;" json:"title"`
	Email     string         `gorm:"type:varchar(255);column:email;" json:"email"`
	DeletedAt gorm.DeletedAt `gorm:"type:timestamp;column:deleted_at" json:"deleted_at,omitempty"`
}

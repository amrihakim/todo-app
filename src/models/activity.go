package models

import (
	"time"

	"gorm.io/gorm"
)

type Activity struct {
	ID        uint64         `gorm:"autoIncrement;primaryKey;column:id" json:"id"`
	Email     string         `gorm:"type:varchar(255);column:email;" json:"email"`
	Title     string         `gorm:"type:varchar(255);column:title;not null;" json:"title"`
	CreatedAt time.Time      `gorm:"type:timestamp;column:created_at;default:current_timestamp" json:"created_at"`
	UpdatedAt *time.Time     `gorm:"type:timestamp;column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"type:timestamp;column:deleted_at" json:"deleted_at,omitempty"`
}

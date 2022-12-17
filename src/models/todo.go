package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID              uint64         `gorm:"autoIncrement;primaryKey;column:id" json:"id"`
	ActivityGroupID string         `gorm:"type:varchar(10);column:activity_group_id" json:"activity_group_id"`
	Title           string         `gorm:"type:varchar(255);column:title" json:"title"`
	IsActive        string         `gorm:"type:varchar(1);column:is_active;default:1" json:"is_active"`
	Priority        string         `gorm:"type:varchar(255);column:priority;default:very-high" json:"priority"`
	CreatedAt       time.Time      `gorm:"type:timestamp;column:created_at;default:current_timestamp" json:"created_at"`
	UpdatedAt       *time.Time     `gorm:"type:timestamp;column:updated_at" json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"type:timestamp;column:deleted_at" json:"deleted_at"`
}

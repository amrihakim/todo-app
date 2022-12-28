package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	CreatedAt       time.Time      `gorm:"type:timestamp;column:created_at;default:current_timestamp" json:"created_at"`
	UpdatedAt       *time.Time     `gorm:"type:timestamp;column:updated_at" json:"updated_at"`
	ID              uint64         `gorm:"autoIncrement;primaryKey;column:id" json:"id"`
	ActivityGroupID uint64         `gorm:"type:int;column:activity_group_id" json:"activity_group_id"`
	Title           string         `gorm:"type:varchar(255);column:title" json:"title"`
	IsActive        bool           `gorm:"type:boolean;column:is_active;default:true" json:"is_active"`
	Priority        string         `gorm:"type:varchar(255);column:priority;default:very-high" json:"priority"`
	DeletedAt       gorm.DeletedAt `gorm:"type:timestamp;column:deleted_at" json:"deleted_at"`
}

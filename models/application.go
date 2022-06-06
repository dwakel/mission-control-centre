package models

import "time"

type Application struct {
	Id          int64     `json:"id" gorm:"column:id"`
	Description string    `json:"description" gorm:"column:description"`
	IsActive    bool      `json:"isActive" gorm:"column:is_active"`
	CreatedAt   time.Time `json:"CreatedAt" gorm:"column:created_at"`
}

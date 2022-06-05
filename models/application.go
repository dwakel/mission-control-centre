package models

type Application struct {
	Id          int64  `json:"id"`
	Description string `json:"description"`
	IsActive    bool   `json:"isActive"`
}

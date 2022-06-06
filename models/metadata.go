package models

import "time"

type Metadata struct {
	Id                   int32     `json:"id" gorm:"column:id"`
	Name                 string    `json:"name" gorm:"column:name"`
	Owner                string    `json:"owner" gorm:"column:owner"`
	ConfigurationManager string    `json:"configurationManager" gorm:"column:configuration_manager"`
	ApplicationId        int32     `json:"applicationId" gorm:"column:application_id"`
	UpdatedAt            time.Time `json:"updatedAt" gorm:"column:updated_at"`
	CreatedAt            time.Time `json:"createdAt" gorm:"column:created_at"`
}

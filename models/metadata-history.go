package models

import "time"

type MetadataHistory struct {
	Id                   int32     `json:"id" gorm:"column:id"`
	MetadataId           int32     `json:"metadataId" gorm:"column:metadata_id"`
	Name                 string    `json:"name" gorm:"column:name"`
	Owner                string    `json:"owner" gorm:"column:owner"`
	ConfigurationManager string    `json:"configurationManager" gorm:"column:configuration_manager"`
	ApplicationId        int32     `json:"applicationId" gorm:"column:application_id"`
	Version              int32     `json:"version" gorm:"column:version"`
	UpdatedAt            time.Time `json:"updatedAt" gorm:"column:updated_at"`
}

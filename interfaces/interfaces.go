package interfaces

import (
	"gorm.io/gorm"
	"mission-control-center/models"
)

type IApplicationRepository interface {
	Create(model models.Application) error
	Get(id int64) (*models.Application, error)
	List() (*[]models.Application, error)
}

type IMetadataRepository interface {
	Create(model models.Metadata) error
	Update(applicationId int64, model models.Metadata) error
	Get(id int64) (*models.Metadata, error)
	GetByApplicationId(id int64) (*models.Metadata, error)
	List() (*[]models.Metadata, error)

	GetMetadataHistory(applicationId int64, limit int32) (*[]models.MetadataHistory, error)

	AddMetadataHistory(applicationId int64, metadata models.Metadata) error

	BeginTransaction() *gorm.DB
}

type IMetadataService interface {
	UpdateTransaction(applicationId int64, model models.Metadata) error
}

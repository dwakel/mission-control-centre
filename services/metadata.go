package services

import (
	"log"
	"mission-control-center/interfaces"
	"mission-control-center/models"
)

type MetadataService struct {
	logger *log.Logger
	repo   *interfaces.IMetadataRepository
}

func NewMetadataService(logger *log.Logger, repo *interfaces.IMetadataRepository) interfaces.IMetadataService {
	return &MetadataService{logger, repo}
}

func (this *MetadataService) UpdateTransaction(metadatId int64, model models.Metadata) error {
	transaction := (*this.repo).BeginTransaction()

	err := (*this.repo).Update(metadatId, model)
	if err != nil {
		return err
	}

	err = (*this.repo).AddMetadataHistory(metadatId, model)
	if err != nil {
		transaction.Rollback()
		return err
	}
	transaction.Commit()
	return nil
}

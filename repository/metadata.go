package repository

import (
	"database/sql"
	"errors"
	"gorm.io/gorm"
	"log"
	"mission-control-center/interfaces"
	"mission-control-center/models"
)

type MetadataRepository struct {
	logger *log.Logger
	repo   *gorm.DB
}

func NewMetadataRepository(logger *log.Logger, repo *gorm.DB) interfaces.IMetadataRepository {
	return &MetadataRepository{logger, repo}
}

func (this *MetadataRepository) Create(model models.Metadata) error {
	command := `INSERT INTO public.metadata 
				(
					name, owner, configuration_manager, application_id
				)
				VALUES
				(@name,@owner,@configurationManager,@applicationId);`
	err := this.repo.Exec(command, map[string]interface{}{"name": model.Name, "owner": model.Owner,
		"configurationManager": model.ConfigurationManager, "applicationId": model.ApplicationId}).Error
	//err := this.repo.Select("description", "is_active").Create(&model).Error
	if err != nil {
		this.logger.Println("failed to add new project")
		this.logger.Println(err)
		return errors.New("failed to add new project")
	}
	this.logger.Println("Successfully inserted new project into database")
	return nil
}

func (this *MetadataRepository) Get(id int64) (*models.Metadata, error) {
	var app models.Metadata
	command := `SELECT m.id, m.application_id, m.name, m.owner,
					m.configuration_manager, m.updated_at, m.created_at
				FROM public.metadata m
					WHERE m.id = @id;`
	err := this.repo.Raw(command, sql.Named("id", id)).Scan(&app).Error
	if err != nil {
		this.logger.Println("failed to add new project")
		this.logger.Println(err)
		return nil, errors.New("failed to add new project")
	}
	this.logger.Println("Successfully fetched application")
	return &app, nil
}

func (this *MetadataRepository) GetByApplicationId(id int64) (*models.Metadata, error) {
	var app models.Metadata
	command := `SELECT m.id, m.application_id, m.name, m.owner,
					m.configuration_manager, m.updated_at, m.created_at
				FROM public.metadata m
					WHERE m.application_id = @id;`
	err := this.repo.Raw(command, sql.Named("id", id)).Scan(&app).Error
	if err != nil {
		this.logger.Println("failed to add new project")
		this.logger.Println(err)
		return nil, errors.New("failed to add new project")
	}
	this.logger.Println("Successfully fetched application")
	return &app, nil
}

func (this *MetadataRepository) List() (*[]models.Metadata, error) {
	var metadata []models.Metadata
	command := `SELECT m.id, m.application_id, m.name, m.owner,
					m.configuration_manager, m.updated_at, m.created_at
				FROM public.metadata m`
	err := this.repo.Raw(command).Scan(&metadata).Error
	if err != nil {
		this.logger.Println("failed to fetch applications")
		this.logger.Println(err)
		return nil, errors.New("failed to fetch applications")
	}
	this.logger.Println("Successfully fetch application list")
	return &metadata, nil
}

func (this *MetadataRepository) Update(metadataId int64, metadata models.Metadata) error {
	command := `UPDATE public.metadata m
				set name = @name,
				owner = @owner,
				configuration_manager = @configurationManager
				WHERE m.id = @metadataId;`
	err := this.repo.Exec(command, map[string]interface{}{"name": metadata.Name, "owner": metadata.Owner,
		"configurationManager": metadata.ConfigurationManager, "metadataId": metadataId}).Error
	if err != nil {
		this.logger.Println("failed to add new project")
		this.logger.Println(err)
		return errors.New("failed to add new project")
	}
	this.logger.Println("Successfully inserted new project into database")
	return nil
}

func (this *MetadataRepository) GetMetadataHistory(metadataId int64, limit int32) (*[]models.MetadataHistory, error) {
	var history []models.MetadataHistory
	command := `SELECT m.id, m.application_id, m.version, m.owner,
					m.configuration_manager, m.updated_at
				FROM public.metadata_history m
				WHERE m.metadata_id = @id
				ORDER BY m.version ACS
				LIMIT @limit`
	err := this.repo.Raw(command, sql.Named("id", metadataId), sql.Named("limit", limit)).Scan(&history).Error
	if err != nil {
		this.logger.Println("failed to fetch applications")
		this.logger.Println(err)
		return nil, errors.New("failed to fetch applications")
	}
	this.logger.Println("Successfully fetch application list")
	return &history, nil
}

func (this *MetadataRepository) AddMetadataHistory(metadataId int64, metadata models.Metadata) error {

	command := `INSERT INTO public.metadata_history(
				application_id, metadata_id, name, owner, version, configuration_manager)
				VALUES (@applicationId, @metadataId, @name, @owner, @version, @configurationManager);`
	err := this.repo.Exec(command, map[string]interface{}{
		"name":                 metadata.Name,
		"owner":                metadata.Owner,
		"configurationManager": metadata.ConfigurationManager,
		"applicationId":        metadata.ApplicationId,
		"metadataId":           metadataId,
		"version":              this.getLatestVersionNumber(int64(metadata.ApplicationId)) + 1}).Error
	if err != nil {
		this.logger.Println("failed to add new project")
		this.logger.Println(err)
		return errors.New("failed to add new project")
	}
	this.logger.Println("Successfully inserted new project into database")
	return nil
}

func (this *MetadataRepository) getLatestVersionNumber(applicationId int64) int64 {
	var version int64
	command := `SELECT m.version
				FROM public.metadata_history m
				WHERE m.application_id = @id
				ORDER BY m.version DESC
				LIMIT 1;`
	_ = this.repo.Raw(command, sql.Named("id", applicationId)).Scan(&version).Error
	return version
}

//This function handles transactions outside the scope of the repository to ensure separations of concerns
func (this *MetadataRepository) BeginTransaction() *gorm.DB {
	return this.repo.Begin()
}

package repository

import (
	"database/sql"
	"errors"
	"gorm.io/gorm"
	"log"
	"mission-control-center/interfaces"
	"mission-control-center/models"
)

type ApplicationRepository struct {
	logger *log.Logger
	repo   *gorm.DB
}

func NewApplicationRepo(logger *log.Logger, repo *gorm.DB) interfaces.IApplicationRepository {
	return &ApplicationRepository{logger, repo}
}

func (this *ApplicationRepository) Create(model models.Application) error {
	command := `INSERT INTO public.applications 
				(
					description, is_active
				)
				VALUES
				(@description,@isActive);`
	err := this.repo.Exec(command, map[string]interface{}{"description": model.Description, "isActive": model.IsActive}).Error
	//err := this.repo.Select("description", "is_active").Create(&model).Error
	if err != nil {
		this.logger.Println("failed to add new project")
		this.logger.Println(err)
		return errors.New("failed to add new project")
	}
	this.logger.Println("Successfully inserted new project into database")
	return nil
}

func (this *ApplicationRepository) Get(id int64) (*models.Application, error) {
	var app models.Application
	command := `SELECT
					a.id, a.description, a.is_active, a.created_at
					FROM public.applications a
					WHERE a.id = @id;`
	err := this.repo.Raw(command, sql.Named("id", id)).Scan(&app).Error
	if err != nil {
		this.logger.Println("failed to add new project")
		this.logger.Println(err)
		return nil, errors.New("failed to add new project")
	}
	this.logger.Println("Successfully fetched application")
	return &app, nil
}

func (this *ApplicationRepository) List() (*[]models.Application, error) {
	var app []models.Application
	command := `SELECT
					a.id, a.description, a.is_active, a.created_at
					FROM public.applications a;`
	err := this.repo.Raw(command).Scan(&app).Error
	if err != nil {
		this.logger.Println("failed to fetch applications")
		this.logger.Println(err)
		return nil, errors.New("failed to fetch applications")
	}
	this.logger.Println("Successfully fetch application list")
	return &app, nil
}

//func (this *ApplicationRepository) GetMetadataHistory(applicationId int64, startDate time.Time, endDate time.Time, limit int64) (models.MetadataHistory, error) {
//
//}

//func (this *ApplicationRepository) AddMetadataHistory(applicationId int64, metadata models.MetadataHistory) (bool, error) {
//
//}

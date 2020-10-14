package database

import (
	"errors"
)

type RepositoryController interface {
	DatabaseCreate(service *Services) error
	DatabaseGet(service *Services) error
	DatabaseSave(service *Services) error
	DatabaseDelete(service *Services) error
}

type MonitorRepo struct{}

var repository RepositoryController

func SetRepoController(repoType RepositoryController) {
	repository = repoType
}

func GetRepoController() RepositoryController {
	return repository
}

func (rp *MonitorRepo) DatabaseCreate(service *Services) error {
	return DB.Create(&service).Error
}

func (rp *MonitorRepo) DatabaseGet(service *Services) error {
	answer := DB.First(&service)
	if answer.RowsAffected == 0 {
		return errors.New("Service not found")
	}
	return answer.Error
}

func (rp *MonitorRepo) DatabaseSave(service *Services) error {
	return DB.Save(&service).Error
}

func (rp *MonitorRepo) DatabaseDelete(service *Services) error {
	answer := DB.Delete(&service)
	if answer.RowsAffected == 0 {
		return errors.New("Url not found")
	}
	return answer.Error
}

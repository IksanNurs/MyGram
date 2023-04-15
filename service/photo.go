package service

import (
	"errors"
	"finalproject_mygram/models"
	"finalproject_mygram/repository"
)

type PService interface {
	GetOnePhoto(id uint) (*models.Photo, error)
	GetAllPhoto() (*[]models.Photo, error)
}

type PhotoService struct {
	Repository repository.PhotoRepository
}

func NewPhotoService(repository repository.PhotoRepository) *PhotoService {
	return &PhotoService{repository}
}

func (service *PhotoService) GetOnePhoto(id uint) (*models.Photo, error) {
	photo := service.Repository.FindById(id)
	if photo == nil {
		return photo, errors.New("photo not found")
	}
	return photo, nil
}

func (service *PhotoService) GetAllPhoto() (*[]models.Photo, error) {
	photo := service.Repository.FindAll()
	if photo == nil || len(*photo)==0{
		return photo, errors.New("data photo not available")
	}
	return photo, nil
}

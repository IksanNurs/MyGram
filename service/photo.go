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
	comment := service.Repository.FindById(id)
	if comment == nil {
		return comment, errors.New("comment not found")
	}
	return comment, nil
}

func (service *PhotoService) GetAllPhoto() (*[]models.Photo, error) {
	comment := service.Repository.FindAll()
	if comment == nil {
		return comment, errors.New("comment not found")
	}
	return comment, nil
}

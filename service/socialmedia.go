package service

import (
	"errors"
	"finalproject_mygram/models"
	"finalproject_mygram/repository"
)

type SService interface {
	GetOneSocialMedia(id uint) (*models.SocialMedia, error)
	GetAllSocialMedia() (*[]models.SocialMedia, error)
}

type SocialMediaService struct {
	Repository repository.SocialMediaRepository
}

func NewSocialMediaService(repository repository.SocialMediaRepository) *SocialMediaService {
	return &SocialMediaService{repository}
}

func (service *SocialMediaService) GetOneSocialMedia(id uint) (*models.SocialMedia, error) {
	comment := service.Repository.FindById(id)
	if comment == nil {
		return comment, errors.New("comment not found")
	}
	return comment, nil
}

func (service *SocialMediaService) GetAllSocialMedia() (*[]models.SocialMedia, error) {
	comment := service.Repository.FindAll()
	if comment == nil {
		return comment, errors.New("comment not found")
	}
	return comment, nil
}

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
	socialmedia := service.Repository.FindById(id)
	if socialmedia == nil {
		return socialmedia, errors.New("socialmedia not found")
	}
	return socialmedia, nil
}

func (service *SocialMediaService) GetAllSocialMedia() (*[]models.SocialMedia, error) {
	socialmedia := service.Repository.FindAll()
	if socialmedia == nil || len(*socialmedia)==0 {
		return socialmedia, errors.New("data socialmedia not available")
	}
	return socialmedia, nil
}

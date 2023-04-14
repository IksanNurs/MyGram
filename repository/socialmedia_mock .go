package repository

import (
	"finalproject_mygram/models"

	"github.com/stretchr/testify/mock"
)

type SocialMediaRepositoryMock struct {
	Mock mock.Mock
}

func (repository *SocialMediaRepositoryMock) FindById(id uint) *models.SocialMedia {
	arguments := repository.Mock.Called(id)

	if arguments.Get(0) == nil {
		return nil
	}
	product := arguments.Get(0).(models.SocialMedia)
	return &product
}


func (repository *SocialMediaRepositoryMock) FindAll() *[]models.SocialMedia {
	arguments := repository.Mock.Called()

	if arguments.Get(0) == nil {
		return nil
	}
	product := arguments.Get(0).([]models.SocialMedia)
	return &product
}

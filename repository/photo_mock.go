package repository

import (
	"finalproject_mygram/models"

	"github.com/stretchr/testify/mock"
)

type PhotoRepositoryMock struct {
	Mock mock.Mock
}

func (repository *PhotoRepositoryMock) FindById(id uint) *models.Photo {
	arguments := repository.Mock.Called(id)

	if arguments.Get(0) == nil {
		return nil
	}
	product := arguments.Get(0).(models.Photo)
	return &product
}


func (repository *PhotoRepositoryMock) FindAll() *[]models.Photo {
	arguments := repository.Mock.Called()

	if arguments.Get(0) == nil {
		return nil
	}
	product := arguments.Get(0).([]models.Photo)
	return &product
}

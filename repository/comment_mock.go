package repository

import (
	"finalproject_mygram/models"

	"github.com/stretchr/testify/mock"
)

type CommentRepositoryMock struct {
	Mock mock.Mock
}

func (repository *CommentRepositoryMock) FindById(id uint) *models.Comment {
	arguments := repository.Mock.Called(id)

	if arguments.Get(0) == nil {
		return nil
	}
	product := arguments.Get(0).(models.Comment)
	return &product
}


func (repository *CommentRepositoryMock) FindAll() *[]models.Comment {
	arguments := repository.Mock.Called()

	if arguments.Get(0) == nil {
		return nil
	}
	product := arguments.Get(0).([]models.Comment)
	return &product
}

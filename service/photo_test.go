package service

import (
	"finalproject_mygram/models"
	"finalproject_mygram/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var photoRepository = &repository.PhotoRepositoryMock{Mock: mock.Mock{}}
var photoService = PhotoService{Repository: photoRepository}

func TestPhotoServiceGetOnePhotoNotFound(t *testing.T) {
	photoRepository.Mock.On("FindById", uint(1)).Return(nil)

	photo, err := photoService.GetOnePhoto(uint(1))

	assert.Nil(t, photo)
	assert.NotNil(t, err)
	assert.Equal(t, "photo not found", err.Error(), "Error response has to be 'photo not found'")
}

func TestPhotoServiceGetOnePhoto(t *testing.T) {
	photo := models.Photo{
		GormModel: models.GormModel{
			ID: 2,
		},
		UserID: 1,
	}

	photoRepository.Mock.On("FindById", uint(2)).Return(photo)

	result, err := photoService.GetOnePhoto(uint(2))

	assert.Nil(t, err)
	assert.NotNil(t, result)

	assert.Equal(t, photo.GormModel.ID, result.GormModel.ID, "Result has to be '2'")
	//assert.Equal(t, photo.Title, result.Title, "Result has to be 'Kaca Mata'")
	assert.Equal(t, &photo, result, "Result has to be photo data with id '2'")
}

func TestPhotoServiceGetAllPhotoNotFound(t *testing.T) {
	photoRepository.Mock.On("FindAll").Return(nil)

	photo, err := photoService.GetAllPhoto()

	assert.Nil(t, photo)
	assert.NotNil(t, err)
	assert.Equal(t, "photo not found", err.Error(), "Error response has to be 'photo not found'")
}

func TestPhotoServiceGetAllPhoto(t *testing.T) {
	photo := []models.Photo{
		{
			GormModel: models.GormModel{
				ID: 1,
			},
			UserID: 2,
		},

		{
			GormModel: models.GormModel{
				ID: 2,
			},
			UserID: 2,
		},
	}
	photoRepository.Mock.On("FindAll").Return(photo)
	result, err := photoService.GetAllPhoto()
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, len(photo), len(*result), "Result lenght")
	assert.Equal(t, photo, *result, "Result value")
}

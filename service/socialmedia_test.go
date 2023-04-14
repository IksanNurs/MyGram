package service

import (
	"finalproject_mygram/models"
	"finalproject_mygram/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var socialmediaRepository = &repository.SocialMediaRepositoryMock{Mock: mock.Mock{}}
var socialmediaService = SocialMediaService{Repository: socialmediaRepository}

func TestSocialMediaServiceGetOneSocialMediaNotFound(t *testing.T) {
	socialmediaRepository.Mock.On("FindById", uint(1)).Return(nil)

	socialmedia, err := socialmediaService.GetOneSocialMedia(uint(1))

	assert.Nil(t, socialmedia)
	assert.NotNil(t, err)
	assert.Equal(t, "socialmedia not found", err.Error(), "Error response has to be 'socialmedia not found'")
}

func TestSocialMediaServiceGetOneSocialMedia(t *testing.T) {
	socialmedia := models.SocialMedia{
		GormModel: models.GormModel{
			ID: 2,
		},
		UserID: 1,
		InputSocialMedia: models.InputSocialMedia{
			Name:           "IG",
			SocialMediaUrl: "https://instagram.com/ramhat",
		},
	}

	socialmediaRepository.Mock.On("FindById", uint(2)).Return(socialmedia)

	result, err := socialmediaService.GetOneSocialMedia(uint(2))

	assert.Nil(t, err)
	assert.NotNil(t, result)

	assert.Equal(t, socialmedia.GormModel.ID, result.GormModel.ID, "Result has to be '2'")
	//assert.Equal(t, socialmedia.Title, result.Title, "Result has to be 'Kaca Mata'")
	assert.Equal(t, &socialmedia, result, "Result has to be socialmedia data with id '2'")
}

func TestSocialMediaServiceGetAllSocialMediaNotFound(t *testing.T) {
	socialmediaRepository.Mock.On("FindAll").Return(nil)

	socialmedia, err := socialmediaService.GetAllSocialMedia()

	assert.Nil(t, socialmedia)
	assert.NotNil(t, err)
	assert.Equal(t, "socialmedia not found", err.Error(), "Error response has to be 'socialmedia not found'")
}

func TestSocialMediaServiceGetAllSocialMedia(t *testing.T) {
	socialmedia := []models.SocialMedia{
		{
			GormModel: models.GormModel{
				ID: 1,
			},
			UserID: 2,
			InputSocialMedia: models.InputSocialMedia{
				Name:           "IG",
				SocialMediaUrl: "https://instagram.com/ramhat",
			},
		},

		{
			GormModel: models.GormModel{
				ID: 2,
			},
			UserID: 2,
			InputSocialMedia: models.InputSocialMedia{
				Name:           "IG",
				SocialMediaUrl: "https://instagram.com/budi",
			},
		},
	}
	socialmediaRepository.Mock.On("FindAll").Return(socialmedia)
	result, err := socialmediaService.GetAllSocialMedia()
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, len(socialmedia), len(*result), "Result lenght")
	assert.Equal(t, socialmedia, *result, "Result value")
}

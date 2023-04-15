package repository

import (
	"finalproject_mygram/models"

	"gorm.io/gorm"
)

type SocialMediaRepository interface {
	FindById(id uint) *models.SocialMedia
	FindAll() *[]models.SocialMedia
}

type SRepository struct {
	db *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) *SRepository {
	return &SRepository{db}
}

func (re *SRepository) FindById(id uint) *models.SocialMedia {
	socialMedia := models.SocialMedia{}
	err:=re.db.Debug().Joins("User").First(&socialMedia, id).Error
	if err!=nil{
		return nil
	}
	return &socialMedia
}

func (re *SRepository) FindAll() *[]models.SocialMedia {
	socialMedia := []models.SocialMedia{}
	err:=re.db.Debug().Joins("User").Order("id ASC").Find(&socialMedia).Error
		if err!=nil{
		return nil
	}
	return &socialMedia
}

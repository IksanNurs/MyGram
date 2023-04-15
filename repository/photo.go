package repository

import (
	"finalproject_mygram/models"

	"gorm.io/gorm"
)

type PhotoRepository interface {
	FindById(id uint) *models.Photo
	FindAll() *[]models.Photo
}

type PRepository struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) *PRepository {
	return &PRepository{db}
}

func (re *PRepository) FindById(id uint) *models.Photo {
	photo := models.Photo{}
	err:=re.db.Debug().Joins("User").First(&photo, id).Error
	if err!=nil{
		return nil
	}
	return &photo
}

func (re *PRepository) FindAll() *[]models.Photo {
	photo := []models.Photo{}
	err:=re.db.Debug().Joins("User").Order("id ASC").Find(&photo).Error
	if err!=nil{
		return nil
	}
	return &photo
}

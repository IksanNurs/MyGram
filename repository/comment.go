package repository

import (
	"finalproject_mygram/models"

	"gorm.io/gorm"
)

type CommentRepository interface {
	FindById(id uint) *models.Comment
	FindAll() *[]models.Comment
}

type CRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CRepository {
	return &CRepository{db}
}

func (re *CRepository) FindById(id uint) *models.Comment {
	comment := models.Comment{}
	err := re.db.Debug().Joins("User").Joins("Photo").First(&comment, id).Error
	if err != nil {
		return nil
	}
	return &comment
}

func (re *CRepository) FindAll() *[]models.Comment {
	comment := []models.Comment{}
	err := re.db.Debug().Joins("User").Joins("Photo").Order("id ASC").Find(&comment).Error
	if err != nil {
		return nil
	}
	return &comment
}

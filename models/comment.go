package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// Comment  represents the model for an Comment
type Comment struct {
	GormModel
	InputComment
	UserID uint `gorm:"not null" json:"user_id" form:"user_id"`
	User   *User
	Photo   *Photo
}

type InputComment struct {
	Message string `gorm:"not null" json:"message" valid:"required~Your message is required"`
	PhotoID uint   `gorm:"not null" json:"photo_id" form:"photo_id"`
}

func (p *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}

func (p *Comment) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errUpdate := govalidator.ValidateStruct(p)

	if errUpdate != nil {
		err = errUpdate
		return
	}
	err = nil
	return
}

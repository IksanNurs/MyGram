package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	GormModel
	InputPhoto
	UserID uint `gorm:"not null" json:"user_id" form:"user_id"`
	User   *User
}

type InputPhoto struct {
	Title    string `gorm:"not null" json:"title" valid:"required~Your title is required"`
	Caption  *string `gorm:"default:null" json:"caption,omitempty" form:"caption"`
	PhotoUrl string `gorm:"not null" json:"photo_url" form:"photo_url" valid:"required~Your photo url is required"`
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}

func (p *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errUpdate := govalidator.ValidateStruct(p)

	if errUpdate != nil {
		err = errUpdate
		return
	}
	err = nil
	return
}

package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	GormModel
	InputSocialMedia
	UserID uint `gorm:"not null" json:"user_id" form:"user_id"`
	User   *User
}

type InputSocialMedia struct {
	Name           string `gorm:"not null" json:"name" valid:"required~Your name is required"`
	SocialMediaUrl string `gorm:"not null" json:"social_media_url" form:"social_media_url" valid:"required~Your social media url is required"`
}

func (p *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}

func (p *SocialMedia) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errUpdate := govalidator.ValidateStruct(p)

	if errUpdate != nil {
		err = errUpdate
		return
	}
	err = nil
	return
}

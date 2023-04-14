package models

import (
	"finalproject_mygram/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	UserName     string        `gorm:"not null;uniqueIndex" json:"full_name" valid:"required~Your user name is required"`
	Email        string        `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Your email is required,email~Invalid email format"`
	Password     string        `gorm:"not null" json:"-" form:"password" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum length of 6 charackters"`
	Age          uint          `gorm:"not null;" json:"age" form:"age" valid:"required~Your age is required,range(9|200)~Age minimum age has a value above 8"`
	Comments     []Comment     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comments"`
	Photos       []Photo       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"photo"`
	SocialMedias []SocialMedia `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"social_media"`
}

type InputUser struct {
	UserName string `gorm:"not null;uniqueIndex" json:"full_name" valid:"required~Your user name is required"`
	Email    string `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Your email is required,email~Invalid email format"`
	Password string `gorm:"not null" json:"password" form:"password" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum length of 6 charackters"`
	Age      uint   `gorm:"not null;" json:"age" form:"age" valid:"required~Your age is required,range(9|200)~Age minimum age has a value above 8"`
}

type LoginUser struct {
	Email    string `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Your email is required,email~Invalid email format"`
	Password string `gorm:"not null" json:"password" form:"password" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum length of 6 charackters"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}
	u.Password = helpers.HassPass(u.Password)
	err = nil
	return

}

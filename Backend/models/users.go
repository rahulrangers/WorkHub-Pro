package models

import "github.com/jinzhu/gorm"

type UserData struct {
	gorm.Model
	Username string `gorm:"unique_index;not null" json:"username"`
	Password string `json:"password"`
	Email    string `gorm:"unique_index;not null" json:"email"`
	WorkHub WorkHub `gorm:"foreignkey:UserDataID" json:"work_hub"`
}

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

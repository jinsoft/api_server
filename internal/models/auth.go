package models

import (
	"gorm.io/gorm"
)

type Auth struct {
	*Model
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
}

func (a Auth) TableName() string {
	return "auth"
}

func (a Auth) Get(db *gorm.DB) (Auth, error) {
	auth := Auth{}
	db = db.Where("app_key =? AND app_secret = ? AND is_delete =?", a.AppKey, a.AppSecret, 0)
	err := db.First(&auth).Error
	if err != nil {
		return auth, err
	}
	return auth, nil
}

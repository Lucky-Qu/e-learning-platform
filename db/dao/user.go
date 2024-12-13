package dao

import "gorm.io/gorm"

var UserDB struct {
	db *gorm.DB
}

func NewUserDB(db *gorm.DB) {
	UserDB.db = db
}

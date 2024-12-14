package dao

import (
	"e-learning-platform/db/model"
	"errors"
	"gorm.io/gorm"
)

type userDB struct {
	db *gorm.DB
}

// UserDB 对外的用户操作接口
var (
	UserDB userDB
)

func initUserDB(db *gorm.DB) {
	UserDB.db = db
	err := UserDB.db.AutoMigrate(&model.User{})
	if err != nil {
		panic(errors.New("数据库自动建表失败"))
	}
}
func (dao userDB) NewUser(user *model.User) error {
	if err := dao.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

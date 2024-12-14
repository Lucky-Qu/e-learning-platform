package service

import (
	"e-learning-platform/db/dao"
	"e-learning-platform/db/model"
	"errors"
)

func StudentRegister(student *model.User) error {
	//用户身份必须为学生
	if student.Identity != "student" {
		return errors.New("传入身份有误")
	}
	//用户名长度不得高于16位
	if len(student.Username) > 16 {
		return errors.New("用户名长度过长")
	}
	//用户密码不得低于8位高于16位
	if len(student.Password) > 16 || len(student.Password) < 8 {
		return errors.New("密码不得低于8位或高于16位")
	}
	if err := dao.UserDB.NewUser(student); err != nil {
		return err
	}
	return nil
}

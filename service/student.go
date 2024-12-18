package service

import (
	"e-learning-platform/cache/redis"
	"e-learning-platform/db/dao"
	"e-learning-platform/db/model"
	"e-learning-platform/log/logger"
	"e-learning-platform/package/util/jwt"
	"errors"
	"go.uber.org/zap"
	"time"
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

func StudentLogin(student *model.User) (string, error) {
	//登陆页面必须为学生登陆页面
	if student.Identity != "student" {
		return "", errors.New("传入身份有误")
	}
	//用户名长度不得高于16位
	if len(student.Username) > 16 {
		return "", errors.New("用户名长度过长")
	}
	//用户密码不得低于8位高于16位
	if len(student.Password) > 16 || len(student.Password) < 8 {
		return "", errors.New("密码不得低于8位或高于16位")
	}
	user, err := dao.UserDB.FindUserByUsername(student.Username)
	if err != nil {
		return "", err
	}
	if student.Identity != user.Identity {
		return "", errors.New("非学生用户")
	}
	if user.Password != student.Password {
		return "", errors.New("密码不正确！")
	}
	tokenString, err := redis.GetStringFromRedis(user.Username)
	if err == nil {
		return tokenString, nil
	}
	tokenString, err = jwt.GetTokenString(user.Username, user.Identity)
	if err != nil {
		return "", err
	}
	err = redis.AddStringToRedis(user.Username, tokenString, time.Hour*1)
	if err != nil {
		logger.DefaultLogger.Logger.Warn("Redis缓存写入失败", zap.Any("err", err))
	}
	return tokenString, nil
}

func StudentUpdate(tokenString string, student *model.User) error {
	//解析Token
	username, password, err := jwt.ParseToken(tokenString)
	if err != nil {
		return errors.New("token解析失败")
	}
	//修改名字
	if username != student.Username {
		//用户名不能重复
		if exist, err := dao.UserDB.UsernameHasExist(username); err != nil {
			return err
		} else if exist {
			return errors.New("用户名不能重复")
		}
		//用户名长度不得高于16位
		if len(student.Username) > 16 {
			return errors.New("用户名长度过长")
		}
	}
	//不能更改身份
	if student.Identity != "student" {
		return errors.New("不能更改身份")
	}
	//修改密码
	if password != student.Password {
		//用户密码不得低于8位高于16位
		if len(student.Password) > 16 || len(student.Password) < 8 {
			return errors.New("密码不得低于8位或高于16位")
		}
	}
	student.Avatar = ""
	err = dao.UserDB.UpdateUser(username, student)
	if err != nil {
		return err
	}
	return nil
}

//TODO 更新头像功能

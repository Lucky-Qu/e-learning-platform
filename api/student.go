package api

import (
	"e-learning-platform/db/model"
	"e-learning-platform/log/logger"
	"e-learning-platform/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func StudentRegister(c *gin.Context) {
	student := model.User{}
	err := c.ShouldBindJSON(&student)
	if err != nil {
		logger.DefaultLogger.Logger.Error("注册失败！", zap.Any("err", err))
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "传入参数有误",
		})
		return
	}
	if err := service.StudentRegister(&student); err != nil {
		logger.DefaultLogger.Logger.Error("数据库注册失败", zap.Any("err", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "注册操作失败",
			"err":  err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "用户注册成功",
		"data": student,
	})
}

func StudentLogin(c *gin.Context) {
	var student model.User
	if err := c.ShouldBindJSON(&student); err != nil {
		logger.DefaultLogger.Logger.Error("绑定登陆参数失败", zap.Any("err", err))
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "传入参数有误",
		})
		return
	}
	tokenString, err := service.StudentLogin(&student)
	if err != nil {
		if err.Error() == "record not found" {
			logger.DefaultLogger.Logger.Info("用户不存在")
			c.JSON(http.StatusNotFound, gin.H{
				"code": http.StatusNotFound,
				"msg":  "用户不存在",
			})
			return
		}
		logger.DefaultLogger.Logger.Error("登陆操作失败", zap.Any("err", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "登陆操作失败，请稍后再试",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"msg":   "登陆成功",
		"token": tokenString,
	})
}

func StudentUpdate(c *gin.Context) {}
func StudentDelete(c *gin.Context) {}

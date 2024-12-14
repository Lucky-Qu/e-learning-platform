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

}

func StudentUpdate(c *gin.Context) {}
func StudentDelete(c *gin.Context) {}

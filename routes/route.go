package routes

import (
	"e-learning-platform/api"
	"e-learning-platform/config"
	"e-learning-platform/middleware"
	"e-learning-platform/middleware/jwt"
	"errors"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.GinLogger())
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "pong",
		})
	})
	// Student用户组
	studentGroup := router.Group("/student")
	{
		//注册学生用户
		studentGroup.POST("/register", api.StudentRegister)
		studentGroup.POST("/login", api.StudentLogin)
		//需要认证操作的行为
		verifiedStudent := studentGroup.Group("/verified")
		verifiedStudent.Use(jwt.IdentifyStudent())
		{
			verifiedStudent.PUT("/update", api.StudentUpdate)
		}
	}
	// Teacher用户组
	teacherGroup := router.Group("/teacher")
	{
		//注册教师用户
		teacherGroup.POST("/register", api.TeacherRegister)
		//需要认证操作的行为
		verifiedTeacher := teacherGroup.Group("/verified")
		verifiedTeacher.Use(jwt.IdentifyTeacher())
		{

		}
	}
	return router
}

func StartService(route *gin.Engine) {
	err := route.Run(config.Config.Server.Host + ":" + config.Config.Server.Port)
	if err != nil {
		panic(errors.New("服务启动失败"))
	}
}

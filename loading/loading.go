package loading

import (
	"e-learning-platform/cache/redis"
	"e-learning-platform/config"
	"e-learning-platform/db/dao"
	"e-learning-platform/log/logger"
)

func Loading() {
	//加载配置
	config.LoadConfig("config/config.json")
	//初始化日志
	logger.InitLogger()
	//初始化数据库
	dao.InitMySQL()
	//初始化Redis
	redis.InitRedis()
}

package config

import (
	"encoding/json"
	"errors"
	"os"
)

type config struct {
	Server struct {
		Host string `json:"host"`
		Port string `json:"port"`
	}
	Mysql struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		Username string `json:"username"`
		Password string `json:"password"`
		Name     string `json:"name"`
	}
	Redis struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		Password string `json:"password"`
	}
	Kafka struct {
		Brokers []string `json:"brokers"`
		Topic   string   `json:"topic"`
	}
	JWT struct {
		Secret string `json:"secret"`
	}
	Log struct {
		Level string `json:"level"`
		Path  string `json:"path"`
	}
}

var Config config

func LoadConfig(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(errors.New("打开配置文件失败"))
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(errors.New("配置文件无法正常关闭"))
		}
	}(file)
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&Config); err != nil {
		panic(errors.New("读取配置文件失败"))
	}
}

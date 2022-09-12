package tools

import (
	"github.com/spf13/viper"
	"os"
	"sreGateway/global"
)

func InitViperConfig() (config *viper.Viper) {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	config = viper.New()
	config.AddConfigPath(path + "/config") //设置读取的文件路径
	config.SetConfigName("config")         //设置读取的文件名
	config.SetConfigType("yaml")           //设置文件的类型
	//尝试进行配置读取
	config.WatchConfig()
	if err := config.ReadInConfig(); err != nil {
		global.LOGGER.Error("Failed to read config file: %s", err.Error())
	}
	global.LOGGER.Info("config change")
	return
}

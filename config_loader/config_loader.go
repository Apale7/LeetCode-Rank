package config

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	configPath = "config/"
	name       = "conf"
)

var rtViper = viper.New()

// 初始化服务端配置
func Init() {
	err := initConfig()
	if err != nil {
		logrus.Fatalf("init config failed, err: %v", err)
		panic(fmt.Errorf("init config err: %w", err))
	}
}

func initConfig() error {
	// 读取基础配置
	rtViper.AddConfigPath(configPath)
	rtViper.AddConfigPath(".")

	// 读取部署环境的配置
	rtViper.SetConfigName(name) // 读取yaml配置文件
	rtViper.AutomaticEnv()
	logrus.Infof("merge conf with %s", name)

	if err := rtViper.MergeInConfig(); err != nil {
		logrus.Fatalf("merge config failed, error: %v", err)
		return err
	}
	return nil
}

// 读取指定名字的配置文件
func read(name string) {
	logrus.Infof("loading conf for %s", name)

	rtViper.SetConfigName(name)

	if err := rtViper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			logrus.Warnf("no such config file")
		} else {
			logrus.Fatalf("read config error")
		}
		panic("load conf file error: " + err.Error())
	}
}

func GetAll() map[string]interface{} {
	return rtViper.AllSettings()
}

func Get(key string) string {
	return rtViper.GetString(key)
}

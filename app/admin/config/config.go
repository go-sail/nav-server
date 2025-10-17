package config

import (
	sailConfig "github.com/keepchen/go-sail/v3/sail/config"
)

type AppConfig struct {
	SailConf sailConfig.Config `yaml:",inline" toml:",inline" json:",inline"` //sail框架配置
	AppName  string            `yaml:"app_name"`                              //应用名称
	Debug    bool              `yaml:"debug"`                                 //是否是debug模式
	Nav      NavConf           `yaml:"nav_conf"`                              //导航配置
}

type NavConf struct {
	RepeatLogin  bool   `yaml:"repeat_login"`  //是否允许重复登录（单态登录）
	IconPath     string `yaml:"icon_path"`     //图标地址目录
	IconEndpoint string `yaml:"icon_endpoint"` //图标访问节点
	InitUser     struct {
		Username string `yaml:"username"` //用户名
		Password string `yaml:"password"` //密码
	} `yaml:"init_user"` //初始化用户
}

var appConfig = &AppConfig{}

func Get() *AppConfig {
	return appConfig
}

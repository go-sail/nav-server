package config

import (
	"fmt"
	"github.com/keepchen/go-sail/v3/lib/etcd"
	"github.com/keepchen/go-sail/v3/sail"
	"gopkg.in/yaml.v2"
	"strings"
)

func parser(configName string, content []byte, viaWatch bool) {
	if !viaWatch {
		fmt.Printf("读取配置文件：%s\n", configName)
	} else {
		fmt.Printf("监听到配置文件：%s 发生变化\n", configName)
	}
	var cfg AppConfig
	err := yaml.Unmarshal(content, &cfg)
	if err != nil {
		panic(err)
	}
	appConfig = &cfg
}

// ParseAndWatchFromFile 解析并监听配置-从文件
func ParseAndWatchFromFile(configName string) {
	sail.Config(true, parser).ViaFile(configName).Parse(parser)
}

// ParseAndWatchFromEtcd 解析并监听配置-从etcd
func ParseAndWatchFromEtcd(endpoints, username, password string, configName string) {
	conf := etcd.Conf{
		Endpoints: strings.Split(endpoints, ","),
		Username:  username,
		Password:  password,
	}
	sail.Config(true, parser).ViaEtcd(conf, configName).Parse(parser)
}

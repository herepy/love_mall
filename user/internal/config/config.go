package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	Wechat struct {
		Appid  string
		Secret string
	}
	Mysql struct {
		Host string
		Port int
		User string
		Pwd  string
		Db   string
	}
	Redis struct {
		Host string
		Type string `json:",default=node,options=node|cluster"`
		Pass string `json:",optional"`
		Tls  bool   `json:",optional"`
	}
	Oss struct {
		Key      string `json:"key"`
		Secret   string `json:"secret"`
		Endpoint string `json:"endpoint"`
		Bucket   string `json:"bucket"`
	}
}

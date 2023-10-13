package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/gorm"
	"love_mall/user/internal/config"
	"love_mall/user/internal/model"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	Redis  *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		DB:     model.GetDB(),
		Redis: redis.MustNewRedis(redis.RedisConf{
			Host: c.Redis.Host,
			Type: c.Redis.Type,
			Pass: c.Redis.Pass,
			Tls:  c.Redis.Tls,
		}),
	}
}

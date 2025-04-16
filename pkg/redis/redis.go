package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"web_go/settings"
)

// Rdb 声明一个全局的rdb变量
var Rdb *redis.Client

// Init 初始化连接
func Init(cfg *settings.RedisConfig) (err error) {
	Rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			cfg.Host,
			cfg.Port),
		Password: "",     // no password set
		DB:       cfg.DB, // use default DB
		PoolSize: cfg.PoolSize,
	})

	_, err = Rdb.Ping().Result()
	if err != nil {
		return err
	}
	zap.L().Info("redis init success")
	return nil
}

func Close() {
	_ = Rdb.Close()
}

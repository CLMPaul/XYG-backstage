package service

import (
	"context"
	"errors"
	"github.com/go-redis/redis"
	"os"
	"strconv"
	"time"
	"xueyigou_demo/cache"

	"xueyigou_demo/constants/redis_keys"
)

var (
	RoleService    roleService
	AccountService accountService
)

var SessionExpiration = time.Hour

var RedisSupportsGetEx bool

func init() {
	// 测试 redis 是否支持 GETEX 命令（redis >= 6.2.0）
	err := cache.GetClient().GetEx(context.Background(), redis_keys.Dymmy(), time.Second).Err()
	if err == nil || errors.Is(err, redis.Nil) {
		RedisSupportsGetEx = true
	}
	if val, ok := os.LookupEnv("SESSION_EXPIRATION"); ok {
		if intVal, err := strconv.ParseInt(val, 10, 64); err == nil {
			SessionExpiration = time.Duration(intVal) * time.Minute
		}
	}
}

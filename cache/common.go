package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"os"
	"strings"
	"time"
	"xueyigou_demo/config"
)

var client redis.UniversalClient

// Redis 在中间件中初始化redis链接
func init() {
	if config.Config.Redis.MasterName != "" {

		// 哨兵模式
		client = redis.NewFailoverClusterClient(&redis.FailoverOptions{
			// The master name.
			MasterName: config.Config.Redis.MasterName,
			// A seed list of host:port addresses of sentinel nodes.
			SentinelAddrs: strings.Split(config.Config.Redis.SentinelAddresses, ","),
			// Sentinel password from "requirepass <password>" (if enabled) in Sentinel configuration
			SentinelPassword: os.Getenv("REDIS_SENTINEL_PASSWORD"),
			// RouteByLatency: true,
			RouteRandomly:   true,
			Username:        config.Config.Redis.UserName,
			Password:        config.Config.Redis.Password,
			PoolSize:        config.Config.Redis.PoolSize,
			MinIdleConns:    config.Config.Redis.MinIdleConnects,
			PoolTimeout:     time.Duration(config.Config.Redis.PoolTimeout) * time.Second,
			ConnMaxIdleTime: time.Duration(config.Config.Redis.IdleTimeout) * time.Second,
		})
	} else {
		adds := strings.Split(config.Config.Redis.Address, ",")
		if len(adds) > 1 { //	集群模式
			client = redis.NewClusterClient(&redis.ClusterOptions{
				Addrs:           adds,
				ReadOnly:        true,
				RouteRandomly:   true,
				PoolSize:        config.Config.Redis.PoolSize,
				MinIdleConns:    config.Config.Redis.MinIdleConnects,
				Username:        config.Config.Redis.UserName,
				Password:        config.Config.Redis.Password,
				PoolTimeout:     time.Duration(config.Config.Redis.PoolTimeout) * time.Second,
				ConnMaxIdleTime: time.Duration(config.Config.Redis.IdleTimeout) * time.Second,
			})
		} else {
			client = redis.NewClient(&redis.Options{
				Addr:            config.Config.Redis.Address,
				Username:        config.Config.Redis.UserName,
				Password:        config.Config.Redis.Password,
				DB:              config.Config.Redis.DB,
				PoolSize:        config.Config.Redis.PoolSize,
				MinIdleConns:    config.Config.Redis.MinIdleConnects,
				PoolTimeout:     time.Duration(config.Config.Redis.PoolTimeout) * time.Second,
				ConnMaxIdleTime: time.Duration(config.Config.Redis.IdleTimeout) * time.Second,
			})
		}
	}
}

func Stop() {
	if client != nil {
		client.Close()
	}
}

func GetClient() redis.UniversalClient {
	return client
}

func Set(key string, value interface{}) error {
	return client.Set(context.Background(), key, value, 0).Err()
}

func SetByTTL(key string, value interface{}, ttl time.Duration) error {
	return client.SetArgs(context.Background(), key, value, redis.SetArgs{TTL: ttl, KeepTTL: true}).Err()
}

func Get(key string) (string, error) {
	return client.Get(context.Background(), key).Result()
}
func GetResult(key string) *redis.StringCmd {
	return client.Get(context.Background(), key)
}

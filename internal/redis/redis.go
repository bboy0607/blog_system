package redis

import (
	"context"
	"membership_system/global"

	"github.com/go-redis/redis/v8"
)

func NewRedisClient() (*redis.Client, error) {
	ctx := context.Background()

	client := redis.NewClient(&redis.Options{
		Addr:     global.RedisSetting.Host + ":" + global.RedisSetting.Port, // 替換成您的Redis伺服器地址和端口
		Password: "",                                                        // 如有需要，設置密碼
		DB:       0,                                                         // 使用默認的數據庫
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}

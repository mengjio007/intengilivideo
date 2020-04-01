package cache

import (
	"fmt"
	"strconv"

	"github.com/go-redis/redis"
)

// RedisClient Redis缓存客户端单例
var RedisClient *redis.Client

// Redis 在中间件中初始化redis链接
func Redis() {
	client := redis.NewClient(&redis.Options{
		Addr:  "47.107.224.190:6379",
		Password: "",
		DB:    1,
	})

	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}

	RedisClient = client
}

const (
	//每日排行
	DailyRankKey = "rank:daily"
)

func VideoViewKey(id uint) string {
	return fmt.Sprintf("view:video:%s", strconv.Itoa(int(id)))
}


func init(){
	Redis()
}

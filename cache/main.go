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
		Addr:  "",
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
func VideoStar(id uint)string{
	return fmt.Sprintf("star:video:%s",strconv.Itoa(int(id)))
}
func VideoLStar(id uint)string{
	return fmt.Sprintf("lstar:video:%s",strconv.Itoa(int(id)))
}
func init(){
	Redis()
}

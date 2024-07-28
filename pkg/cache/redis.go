package cache

import (
    "github.com/go-redis/redis/v8"
    "context"
)

var ctx = context.Background()

func GetRedisClient() *redis.Client {
    rdb := redis.NewClient(&redis.Options{
        Addr: "localhost:6379",
        Password: "", // no password set
        DB: 0,  // use default DB
    })
    return rdb
}

func SetCache(key string, value string) error {
    client := GetRedisClient()
    err := client.Set(ctx, key, value, 0).Err()
    return err
}

func GetCache(key string) (string, error) {
    client := GetRedisClient()
    val, err := client.Get(ctx, key).Result()
    return val, err
}

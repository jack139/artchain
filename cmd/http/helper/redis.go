package helper

import (
	"context"

	"github.com/go-redis/redis/v8"
)

const (
	imageKeyPrefix  = "image_"
)

var (
	rdb *redis.Client
	ctx = context.Background()
)

func InitRDB() error {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:7480",
		Password: "e18ffb7484f4d69c2acb40008471a71c",
		DB:       0,  // use default DB
	})

	if _, err := rdb.Ping(ctx).Result(); err!=nil {
		return err
	}

	return nil
}

func CacheImage(key string, data []byte) error {
	err := rdb.Set(ctx, imageKeyPrefix+key, string(data), 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func GetImage(key string) ([]byte, error) {
	val, err := rdb.Get(ctx, imageKeyPrefix+key).Result()
	if err == redis.Nil{
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return []byte(val), nil
}

func DelImage(key string) error {
	err := rdb.Get(ctx, imageKeyPrefix+key).Err()
	if err != nil {
		return err
	}
	return nil
}

package redis

//проверяет контент на содержание плохих слов и блокирует такой контент.
import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

type blogClient struct {
	rdb *redis.Client
}

func InitClient(ctx context.Context) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// Проверяем подключение
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Не удалось подключиться к Redis: %v", err)
	}

	return rdb
}

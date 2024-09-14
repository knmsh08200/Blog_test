package redis

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/go-redis/redis/v8"
)

type blogClient struct {
	rdb *redis.Client
}

func NewRep(r *redis.Client) *blogClient {
	return &blogClient{rdb: r}
}

func (r *blogClient) СontainsProfanity(ctx context.Context, rdb *redis.Client, content string) bool {
	words, err := rdb.SMembers(ctx, "bad_words_set").Result()
	if err != nil {
		log.Printf("Ошибка при получении списка плохих слов из Redis: %v", err)
		return true
	}

	for _, word := range words {
		if contains(content, word) {
			fmt.Println("Banned words detected")
			return true
		}
	}
	return false
}

func contains(content, word string) bool {
	return strings.Contains(content, word)
}

func (r *blogClient) LoadProfanityWords(ctx context.Context, db *sql.DB) {

	rows, err := db.Query("SELECT words FROM profanitywords ")

	if err != nil {
		log.Fatalf("Bad fetching profinity words from DB %v", err)
	}

	defer rows.Close()
	for rows.Next() {
		var word string
		if err := rows.Scan(&word); err != nil {
			log.Fatalf("Ошибка при сканировании строки: %v", err)
		}

		err := r.rdb.SAdd(ctx, "bad_words_set", word).Err()
		if err != nil {
			log.Fatalf("Не удалось сохранить слово '%s' в Redis: %v", word, err)
		}
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("Ошибка при переборе строк: %v", err)
	}
}

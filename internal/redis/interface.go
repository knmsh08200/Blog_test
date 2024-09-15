package redis

import (
	"context"
	"database/sql"

	"github.com/go-redis/redis/v8"
)

type ClientRep interface {
	LoadProfanityWords(ctx context.Context, db *sql.DB)
}

type ProfanityChecker interface {
	СontainsProfanity(ctx context.Context, rdb *redis.Client, content string) bool
}

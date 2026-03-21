package database

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func ConnectRedis(host string, port string) error {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: host + ":" + port,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return RedisClient.Ping(ctx).Err()
}

func DisconnectRedis() {
	if RedisClient == nil {
		return
	}

	if err := RedisClient.Close(); err != nil {
		log.Printf("Erro ao desconectar Redis: %v", err)
	}
}

func IsRedisConnected() bool {
	if RedisClient == nil {
		return false
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	return RedisClient.Ping(ctx).Err() == nil
}

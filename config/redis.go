package config

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()
var RedisClient *redis.Client

func ConnectRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "54.160.106.5:6379", // o "redis:6379" si estás dentro del contenedor
		Password: "CartService123!",   // contraseña establecida en docker-compose
		DB:       0,                   // base de datos por defecto
	})

	_, err := RedisClient.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("Redis connection failed: %v", err)
	}

	log.Println("SRedis connected successfully")
}

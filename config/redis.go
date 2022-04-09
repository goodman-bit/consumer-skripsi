package config

import (
	"flag"
	"fmt"

	"github.com/go-redis/redis"
)

func ConnectRedisLocal() *redis.Client {
	connStr := fmt.Sprintf("%s:%s", REDISHostLocal, REDISPortLocal)
	var addr = flag.String("Server Local", connStr, "Redis server address")
	fmt.Println("Successful Connected to Redis Local:", string(*addr))

	rdb := redis.NewClient(&redis.Options{
		Addr:     *addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return rdb
}

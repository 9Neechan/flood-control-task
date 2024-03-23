package configs

import(
	"fmt"
	"log"
	"context"

	"github.com/redis/go-redis/v9"
)

func ConnectRedisDB() *redis.Client {
	GetEnv()
	
	client := redis.NewClient(&redis.Options{
        Addr:	  EnvRedisAddr(),
        Password: EnvRedisPassword(), // no password set
        DB:		  EnvRedisDBName(),  // use default DB
    })

	if err := client.Ping(context.TODO()).Err(); err != nil {
        log.Fatal(err)
	}

	fmt.Println("Connected to RedisDB!")

	return client
}


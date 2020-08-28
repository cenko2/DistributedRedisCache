package CacheLogic

import (
	"context"
	"fmt"

	"github.com/go-redis/redis"
)

type RedisCache struct {
	Rdb *redis.Client
}

var ctx = context.Background()

func (m RedisCache) Get(key string) string {
	val, err := m.Rdb.Get(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	return val
}
func (m RedisCache) Insert(key string, val string, ttl int) {
	err := m.Rdb.Set(ctx, key, val, 0).Err()
	if err != nil {
		panic(err)
	}
}
func (m RedisCache) KeyExists(key string) bool {
	_, err := m.Rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		fmt.Println(key, "does not exist")
		return false
	} else if err != nil {
		panic(err)
	} else {
		return true
	}
}

func ExampleNewClient() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:7001",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := rdb.Ping(ctx).Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
}

func ExampleClient() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:7001",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}

package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func ExampleClient() {
	// rdb作为一个全局变量,在服务初始化的时候实例化出来,设置好连接池,以后在这个服务的生命周期之内,直接拿来用即可
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
		PoolSize: 10, // 连接池数量
	})

	err := rdb.Set(ctx, "key", "hello world", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key:", val)

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

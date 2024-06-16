package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	rdb    *redis.Client
	id_cnt int
)

func initDB() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,   // use default DB
		PoolSize: 100, // 连接池大小
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = rdb.Ping(ctx).Result()

	return err
}

func push_in_db(word string, docIDs []string) {
	ctx := context.Background()

	err := rdb.RPush(ctx, word, docIDs).Err()
	if err != nil {
		fmt.Println("insert Redis fail:", err)
		fmt.Println("word: ", word, "  docIDs: ", docIDs)
		return
	}

}

func vis_list(url string) {
	ctx := context.Background()
	val, err := rdb.Get(ctx, "id_cnt").Result()
	if err != nil {
		panic(err)
	}

	id_cnt, err = strconv.Atoi(val)

	err = rdb.Set(ctx, "id_cnt", id_cnt+1, 0).Err()
	if err != nil {
		panic(err)
	}

	err = rdb.Set(ctx, strconv.Itoa(id_cnt), url, 0).Err()
	if err != nil {
		panic(err)
	}
	err = rdb.Set(ctx, url, strconv.Itoa(id_cnt), 0).Err()
	if err != nil {
		panic(err)
	}
}

func title_push_in_db(url string, title string) {
	ctx := context.Background()
	idx := find_key(url)
	err := rdb.Set(ctx, "title"+idx, title, 0).Err()
	if err != nil {
		panic(err)
	}

}

func find_key(key string) string {
	ctx := context.Background()
	value, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		fmt.Println("Key does not exist")
	} else if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("")
		//fmt.Println("\x1b[36m", value, "\x1b[0m")
	}
	return value
}

func list_end(word string) (string, error) {
	ctx := context.Background()
	result, err := rdb.LIndex(ctx, word, -1).Result()
	if err != nil {
		return "", err
	}
	return result, nil
}

func Gets(key string) string {
	ctx := context.Background()
	value, err := rdb.Get(ctx, "4").Result()
	if err == redis.Nil {
		fmt.Println("Key does not exist")
	} else if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Value:", value)
	}
	return value
}

func Get_index_list(key string) []string {
	ctx := context.Background()
	list, err := rdb.LRange(ctx, key, 0, -1).Result()
	if err != nil {
		fmt.Println("读取列表失败:", err)
		panic(err)
	}
	return list
}

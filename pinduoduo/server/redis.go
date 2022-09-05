package server

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"log"
)

var (
	rdb *redis.Client
)

type ReturnCodeResponse struct {
	Code  string `json:"code"`  // 返回code
	State string `json:"state"` // 状态
}

// 初始化连接
func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123456", // no password set
		DB:       0,        // use default DB
		PoolSize: 100,      // 连接池大小
	})

	_, err = rdb.Ping().Result()
	return err
}

func Save(d ReturnCodeResponse) {
	if err := initClient(); err != nil {
		return
	}
	jsonData, _ := json.Marshal(&d)
	err := rdb.Set("ReturnCodeResponse", jsonData, 0).Err()
	if err != nil {
		panic(err)
	}
}

func GetVal() ReturnCodeResponse {
	if err := initClient(); err != nil {
		log.Println("redis get error")
	}
	val, err := rdb.Get("ReturnCodeResponse").Result()
	if err != nil {
		panic(err)
	}
	var v ReturnCodeResponse
	if err == redis.Nil {
		fmt.Println("ReturnCodeResponse does not exist")
	} else if err != nil {
		panic(err)
	}
	_ = json.Unmarshal([]byte(val), &v)
	return v
}

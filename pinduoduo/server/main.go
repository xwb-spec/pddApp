package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
)

var (
	rdb *redis.Client
	a   string
)

type ReturnCodeResponse struct {
	Code  string `json:"code"`  // 返回code
	State string `json:"state"` // 状态
}

type Account struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 初始化连接
func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "fefeflfeinifenifel@11w232", // no password set
		DB:       0,                           // use default DB
		PoolSize: 100,                         // 连接池大小
	})
	_, err = rdb.Ping().Result()
	return err
}

// 保存数据
func StoreJson(k string, v any) error {
	if err := initClient(); err != nil {
		return err
	}
	jsonData, _ := json.Marshal(&v)
	if err := rdb.Set(k, jsonData, 0).Err(); err != nil {
		return err
	}
	return nil
}
func StoreCode(v any) {
	err := StoreJson("ReturnCodeResponse", v)
	if err != nil {
		log.Fatal(err)
	}
}

func LoadKey(key string) (string, error) {
	if err := initClient(); err != nil {
		return "", err
	}
	val, err := rdb.Get(key).Result()
	if err == redis.Nil {
		return "", errors.New(fmt.Sprintf("key %s does not exists", key))
	} else if err != nil {
		return "", err
	}
	return val, nil
}
func LoadCode() ReturnCodeResponse {
	val, err := LoadKey("ReturnCodeResponse")
	if err != nil {
		log.Fatal(err)
	}
	var v ReturnCodeResponse
	_ = json.Unmarshal([]byte(val), &v)
	return v
}

func PwdHash(pwd string) (string, error) {
	// 第二个参数是进行哈希的次数，这里采用了默认值10,数字越大生成的密码速度越慢，成本越大。但是更安全
	// bcrypt每次生成的编码是不同的，较于md5更安全
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), err
}

func StoreAccount() {
	pwdHash, _ := PwdHash("AI3pJ5txofRJvZ8CKSWrig==")
	err := StoreJson("AuthAccount", Account{
		Username: "13714442790",
		Password: pwdHash,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func LoadAccount() Account {
	val, err := LoadKey("AuthAccount")
	if err != nil {
		log.Fatal(err)
	}
	var v Account
	_ = json.Unmarshal([]byte(val), &v)
	return v
}

func RunApi() {
	r := gin.Default()
	r.GET("/api/v1/callback/", func(c *gin.Context) {
		code, ok := c.GetQuery("code")
		if !ok {
			c.JSON(400, "callback code failed")
			return
		}
		state, ok := c.GetQuery("state")
		if !ok {
			c.JSON(400, "callback code failed")
			return
		}
		StoreCode(ReturnCodeResponse{
			Code:  code,
			State: state,
		})
		c.JSON(200, "callback code ok")
		return
	})
	r.POST("/api/v1/callback/", func(c *gin.Context) {
		data := LoadCode()
		c.JSON(200, data)
		return
	})
	r.POST("/api/v1/auth/", func(c *gin.Context) {
		data := LoadAccount()
		c.JSON(200, data)
		return
	})
	r.Run("0.0.0.0:5001")
}

func init() {
	flag.StringVar(&a, "a", "account", "[account|runApi]")
}
func main() {
	flag.Parse()
	switch {
	case a == "account":
		StoreAccount()
	case a == "runApi":
		RunApi()
	default:
		os.Exit(1)
	}
}

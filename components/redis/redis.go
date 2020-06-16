//+build redis  !production

package redis

import (
	"github.com/go-redis/redis"
	"log"
	"x/global"
	"x/utils"
)

type Config struct {
	Host string `map:"host"`
	Port string `map:"port"`
	Auth string `map:"auth"`
	DB   int    `map:"db"`
}

type redisStruct struct {
	DB *redis.Client
}

func (r *redisStruct) NeedConfig() bool {
	return true
}

func RedisConnect(r *Config) *redis.Client {
	redis := redis.NewClient(&redis.Options{
		Addr:     r.Host + ":" + r.Port,
		Password: r.Auth, // no password set
		DB:       r.DB,   // use default DB
	})
	if redis == nil {
		log.Fatalln("redis初始化错误")
	}
	return redis
}

func (r *redisStruct) NewComponent(config interface{}) {
	redisParams := &Config{}
	if conf, ok := config.(map[interface{}]interface{}); ok {
		err := utils.MapToStruct(conf, redisParams)
		if err != nil {
			log.Fatal(err)
		}
		r.DB = RedisConnect(redisParams)
	}else{
		log.Fatalln("redi配置文件错误")
	}

}

var Redis redisStruct

func init() {
	//注册组件
	global.Global.Register("redis", &Redis)
}

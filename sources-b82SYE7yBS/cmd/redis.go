package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

type redisManager struct{
	RedisDb *redis.Client
}
func (r *redisManager)init()(err error){
	r.RedisDb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // redis地址
		Password: "",               // redis密码，没有则留空
		DB:       0,                // 默认数据库，默认是0
	})
	_, err = r.RedisDb.Ping().Result()
	if err != nil {
		fmt.Println("Redis链接错误")
		return err
	}

	return nil
}

func (r* redisManager)get(key string)(str string,err error){
	return r.RedisDb.Get(key).Result()
}
func (r* redisManager)set(key string,val string)(str string,err error){
	return r.RedisDb.Set(key,val,0).Result()
}
func (r* redisManager)exists(key string)(b bool,e error){
	n, err := r.RedisDb.Exists(key).Result()
	if err!=nil{
		return false,err
	}
	return n>0,nil
}


//var RedisManager=redisManager{}

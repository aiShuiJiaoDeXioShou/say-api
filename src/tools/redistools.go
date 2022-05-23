package tools

import (
	"log"
	"time"

	"github.com/garyburd/redigo/redis"
)

func SetRedisVlue(key string, obj string) {
	// 建立连接
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Println("redis.Dial err=", err)
		return
	}
	// 通过go向redis写入数据 string [key - value]
	_, err = conn.Do("Set", key, obj)
	if err != nil {
		log.Println("set err=", err)
		return
	}
	// 关闭连接
	defer conn.Close()
}

func GetRedisVlue(name string) string {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	// 读取数据 获取名字
	r, err := redis.String(conn.Do("Get", name))
	if err != nil {
		log.Println("set err=", err)
		return ""
	}
	defer conn.Close()
	return r
}

func SetTimeVlue(name string) {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	//time.Duration(300)*time.Second这句话的意思是乘以三百分钟
	_,err = conn.Do("expire", name,time.Duration(300)*time.Second)
	if err != nil {
		log.Println("设置过期时间成功！！")
		return
	}
}

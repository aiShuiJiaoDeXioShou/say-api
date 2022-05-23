package test

import (
	"fmt"
	"testing"

	"github.com/garyburd/redigo/redis"
)

func RedisTest1(t *testing.T) {
	// 建立连接
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err=", err)
		return
	}

	// 通过go向redis写入数据 string [key - value]
	_, err = conn.Do("Set", "name", "Tom")
	if err != nil {
		fmt.Println("set err=", err)
		return
	}
	// 关闭连接
	defer conn.Close()

	// 读取数据 获取名字
	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("set err=", err)
		return
	}
	//nameString := r.(string)
	fmt.Println("Manipulate success, the name is", r)
}

/* func TestXxx(t testing.T) {
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	   case i1 = <-c1:
		  fmt.Printf("received ", i1, " from c1\n")
	   case c2 <- i2:
		  fmt.Printf("sent ", i2, " to c2\n")
	   case i3, ok := (<-c3):  // same as: i3, ok := <-c3
		  if ok {
			 fmt.Printf("received ", i3, " from c3\n")
		  } else {
			 fmt.Printf("c3 is closed\n")
		  }
	   default:
		  fmt.Printf("no communication\n")
	} 
} */
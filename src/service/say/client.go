package say

import (
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID     string
	SendID string
	Socket *websocket.Conn
	Send   chan []byte
}

func (c *Client) Read() {
	// 避免忘记关闭，所以要加上close
	defer func() {
		Mang.Unregister <- c
		_ = c.Socket.Close()
	}()

	for {
		c.Socket.PongHandler()
		sendMsg := new(SendMsg)
		err := c.Socket.ReadJSON(&sendMsg) // 读取json格式，如果不是json格式，会报错,把数据读到sendMsg里面
		if err != nil {
			log.Println("数据格式不正确", err)
			Mang.Unregister <- c //退出连接
			_ = c.Socket.Close()
			break
		}
		Mang.Broadcast <- &Broadcast{
			Client:  c,
			Message: []byte(sendMsg.Content),
		}

	}
}
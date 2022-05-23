package test

import (
	"github.com/gorilla/websocket"
)

//发送消息的结构体
type SendMsg struct {
	 Type int `json:"type"`
	 Content string `json:"content"`
}

//接受消息的结构体
type RecvMsg struct {
	From string `json:"from"`
	Code int `json:"code"`
	Content string `json:"content"` 
}

//用户结构体
type Client struct {
	ID string 
	SendID string
	Socket *websocket.Conn
	Send chan []byte
}
//广播结构体包过广播内容和源用户
type BroadCast struct {
	Client *Client
	Message []byte
	Type int
}

//用户管理类
type ClientManager struct {
	Clients 	map[string]*Client
	// Broadcast 	chan *Broadcast
	Reply 		chan *Client
	Register 	chan *Client
	Unregister  chan *Client
}

//信息转JSON (包括：发送者、接收者、内容)
type Message struct {
	Sender 		string 		`json:"sender,omitempty"`
	Recipient 	string 		`json:"recipient,omitempty"`
	Content 	string 		`json:"content,omitempty"`
}




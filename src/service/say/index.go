package say

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var Mang = ManagerClient{
	Clients:    make(map[string]*Client),
	Broadcast:  make(chan *Broadcast),
	Reply:      make(chan *Client),
	Register:   make(chan *Client),
	Unregister: make(chan *Client),
}


//这个是对话服务
func SayHandler(c *gin.Context) {
	id := c.Param("id")
	toid := c.Param("toid")
	log.Println("id:", id, "toid:", toid)

	var upGrader = websocket.Upgrader{  
		CheckOrigin: func (r *http.Request) bool {  
		   return true  
		},  
	 }

	conn, err := upGrader.Upgrade(c.Writer, c.Request, nil) // 升级成ws协议
	
	if err != nil {
		log.Println("升级错误", err)
		http.NotFound(c.Writer, c.Request)
		return
	}

	client := &Client{
		ID:     uid(id, toid),
		SendID: uid(toid, id),
		Socket: conn,
		Send:   make(chan []byte),
	}

	if _cilent := Mang.Clients[client.ID];_cilent != nil {
		log.Println("该用户已上线", _cilent.ID)
		client.Socket.WriteJSON(gin.H{
			"code": 400,
			"data": "该用户已上线",
		})
		client.Socket.Close()
		return
	}

	Mang.Register <- client
	go client.Read()
}



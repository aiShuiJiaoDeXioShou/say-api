package say

import (
	"log"

	"github.com/gin-gonic/gin"
)

type ManagerClient struct {
	Clients    map[string]*Client
	Broadcast  chan *Broadcast
	Reply      chan *Client
	Register   chan *Client
	Unregister chan *Client
}

//启动聊天服务器
func (this *ManagerClient) Start() {
	
	log.Println("ManagerClient Start...")
	for {
		select {
		//每次连接时注册在线用户
		case client := <-this.Register:
			this.Clients[client.ID] = client
			client.Socket.WriteJSON(gin.H{
				"code": 200,
				"data": "连接成功",
			})
			log.Println("用户连接成功", client.ID)

		//注销一个用户
		case client := <-this.Unregister:
			if _, ok := this.Clients[client.ID]; ok {
				delete(this.Clients, client.ID)
				close(client.Send)
				log.Println("用户注销成功", client.ID)
			}

		//广播消息
		case broadcast := <-this.Broadcast:

			//获取将要广播的消息
			message := broadcast.Message
			//获取将要传递的用户
			sendID := broadcast.Client.SendID
			sendClient := this.Clients[sendID]

			//该用户没有上线，将消息先存到缓存里面
			if sendClient == nil {

				broadcast.Client.Socket.WriteJSON(gin.H{
					"type": 1,
					"data": sendID + ":他还没有上线",
				})
				log.Println("用户", broadcast.Client.ID, ",->发送消息给",sendID,"...失败，他还没有上线...该消息为",message)
				break
			}

			//如果不为空就表明,该用户已经上线了,将消息转发给他
			sendClient.Socket.WriteJSON(&SendMsg{
				Type: 1,
				Content: string(message),
			})
		}
	}
}
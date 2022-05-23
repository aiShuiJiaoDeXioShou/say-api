package say


// 广播类，包括广播内容和源用户
type Broadcast struct {
	Client  *Client
	Message []byte
	Type    int
}

// 发送消息的类型
type SendMsg struct {
	Type    int    `json:"type"`
	Content string `json:"content"`
}







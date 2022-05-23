package say

//每一个通讯都有一个user来管理，当用户登录之后创建一个sayuser对象来管理这个用户的所有通讯行为
//SayUser -> Clients -> Client 上线了->
//发送累计的消息 没有累计消息->
//关闭连接 用户发送消息->
//打开连接 判断指定用户是否在线,没有在线关闭连接,储存到持久层去 
//在线->发送消息给指定用户
type SayUser struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Clients []*Client
}

package service

import (
	"encoding/json"
	"github.com/nsqio/go-nsq"
	"net/http"
)

// 用户列表
var SocketList = make(map[int]Client)

// 新建客户端连接
func NewSocketClient(token string, w http.ResponseWriter, r *http.Request) (client *Client) {
	conn, err := upgrader.Upgrade(w, r, w.Header())
	if err != nil {
		return
	}

	client = &Client{
		Conn:  conn,
		Token: token,
	}
	return client
}

// 处理消息
func HandleMessage(msg *nsq.Message) {
	var m Message
	err := json.Unmarshal(msg.Body, &m)
	if err != nil {
		return
	}

	for _, client := range SocketList {
		if client.Id != m.ID { // 自己的消息不发给自己
			client.Conn.WriteJSON(m)
		}
	}
}

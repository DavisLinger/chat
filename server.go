package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"imoniang.com/chat/service"
	"imoniang.com/chat/sql"
	"net/http"
	"time"
)

// WebSocket处理事件
func wsHandler(w http.ResponseWriter, r *http.Request) {

	// 判断是否为WebSocket协议
	if !websocket.IsWebSocketUpgrade(r) {
		http.Error(w, "不是webSocket协议", http.StatusBadRequest)
		return
	}
	token := r.Header.Get("Token") // 获取Token，可通过JWT请求个人数据
	client := service.NewSocketClient(token, w, r)

	// 判断Token，并获取个人信息
	user, result := sql.CheckToken(token)
	if !result {
		client.Conn.Close()
		return
	}

	// 判断此用户是否已登录
	_, ok := service.SocketList[user.ID]
	if ok { // 已登录，通知下线信息
		service.SocketList[user.ID].Conn.WriteJSON(&service.Message{
			ID: -1,
		})
		service.SocketList[user.ID].Conn.Close()
	}

	client.Id = user.ID
	client.Name = user.Nick
	service.SocketList[user.ID] = *client // 将连接信息加入列表中

	defer client.Conn.Close()
	for {
		var m service.Message
		err := client.Conn.ReadJSON(&m)

		if websocket.IsCloseError(err, websocket.CloseNoStatusReceived, websocket.CloseAbnormalClosure, websocket.CloseNormalClosure) {
			delete(service.SocketList, client.Id)
			fmt.Println("用户主动断开了连接")
			return
		}
		if err != nil {
			fmt.Println("读取数据错误", err)
			return
		}
		m.ID = client.Id
		m.SendTime = time.Now().Unix()
		message, err := json.Marshal(m)
		if err != nil {
			fmt.Println("转换数据错误", err)
			return
		}
		SendMessage("Message", message)
		fmt.Printf("获取到数据: %#v\n", m)
	}
}

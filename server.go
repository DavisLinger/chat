package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"imoniang.com/chat/service"
	"net/http"
	"strconv"
	"time"
)

// WebSocket处理事件
func wsHandler(w http.ResponseWriter, r *http.Request) {

	// 判断是否为WebSocket协议
	if !websocket.IsWebSocketUpgrade(r) {
		http.Error(w, "不是webSocket协议", http.StatusBadRequest)
		return
	}
	token := r.Header.Get("Token")            // 获取Token，可通过JWT请求个人数据
	name := r.Header.Get("Name")              // 用户昵称应通过Token来获取
	id, _ := strconv.Atoi(r.Header.Get("Id")) // 此ID应通过Token来获取,此处会发生同ID多地区登录覆盖的问题
	fmt.Printf("有新用户发起连接:Token:%v,Name:%v,Id:%d\n", token, name, id)
	client := service.NewSocketClient(token, name, id, w, r)
	service.SocketList[id] = *client
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

package main

import (
	"log"
	"net/http"
)

func main() {
	// 初始化Nsq生产者
	errorNo, err := InitProducer("127.0.0.1:4150")
	if err != nil {
		switch errorNo {
		case 1:
			log.Fatalf("init producer failed：%v\n", err)
			return
		case 2:
			log.Fatalf("fail to ping %v\n", err)
		}
	}

	// 初始化Nsq消费者
	InitConsumer("Message", "Message-channel", "127.0.0.1:4161")

	// 初始化WebSocket
	http.HandleFunc("/ws", wsHandler)        // 注册Ws路由
	panic(http.ListenAndServe(":8080", nil)) // 设置监听信息
}

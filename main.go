package main

import (
	"imoniang.com/chat/config"
	"imoniang.com/chat/sql"
	"log"
	"net/http"
)

func main() {
	// 初始化Nsq生产者
	errorNo, err := InitProducer(config.ProducerAddr)
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
	InitConsumer("Message", "Message-channel", config.ConsumerAddr)

	sql.InitDb()
	// 初始化WebSocket
	http.HandleFunc("/", index)            // 注册首页路由
	http.HandleFunc("/ws", wsHandler)      // 注册Ws路由
	http.HandleFunc("/login", login)       // 注册登录路由
	http.HandleFunc("/register", register) // 注册注册路由
	http.HandleFunc("/chat", chat)         // 注册聊天路由

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static")))) // 注册静态资源路由
	defer sql.DB.Close()
	// panic(http.ListenAndServeTLS(":8080", "cer/imoniang.crt", "cer/imoniang.key", nil)) // SSL
	panic(http.ListenAndServe(":8080", nil)) // 设置监听信息
}

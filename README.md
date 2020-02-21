# 聊天程序
目前实现的功能为：群聊！

### 用到的第三方包有
- github.com/nsqio/go-nsq
- github.com/gorilla/websocket
- github.com/jinzhu/gorm

## 使用说明
先启动NSQ
```
nsqlookupd
nsqd --lookupd-tcp-address=127.0.0.1:4160
```
然后启动程序
``` 
go run main.go
```
websocket连接 `ws://127.0.0.1:8080/ws`

需要在Header中传递的值有：

| Key        | Value   |  说明  |
| --------   | -----:  | :----:  |
| Token      | 5bfb87a11a3d771c36086afb897e9af3   |   Token，登录后返回的     |

发送消息格式：
`{"message":"我是墨娘"}`

界面：http://127.0.0.1:8080/
注册界面还没对接

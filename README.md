# 墨娘写的第一个项目

####用到的第三方包有
- github.com/nsqio/go-nsq
- github.com/gorilla/websocket

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
| Token      | si82kksu2   |   Token，用来获取数据说明的     |
| Name        |   墨娘   |  用户昵称，为了方便所以放到这里   |
| Id        |    1   |  用户ID，为了方便所以放到这里  |

发送消息格式：
`{"message":"我是墨娘"}`

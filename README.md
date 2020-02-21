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
go build
chat
```
界面：http://127.0.0.1:8080/

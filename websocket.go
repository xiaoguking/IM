package main

import (
	"fmt"
	_ "github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"net/http"
)

type connection struct {
	ws     *websocket.Conn
	sc     chan []byte
	Data   Data
	client string
	token  string
}

var wu = &websocket.Upgrader{ReadBufferSize: 512, WriteBufferSize: 512, CheckOrigin: func(r *http.Request) bool { return true }}

func Handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Server", " IM/1.0")
	ws, err := wu.Upgrade(w, r, w.Header())
	if err != nil {
		return
	}
	c := &connection{sc: make(chan []byte, 521), ws: ws}
	//获取token
	token := r.URL.Query().Get("token")
	if len(token) == 0 {
		c.ws.Close()
		return
	}
	//向全局的client推送消息
	h.r <- c
	//生成一个客户唯一的标识码
	client := UniqueId()
	c.client = client
	//存储在线的client客户端
	clientList[client] = client
	//给每一个客户端绑定一个唯一的标识码
	u.m[client] = c
	//绑定uid 到客户端
	lock.Lock()
	value, ok := uidBindClient[token]
	lock.Unlock()
	if !ok {
		value = make([]string, 0, 1)
	}
	value = append(value, client)
	lock.Lock()
	uidBindClient[token] = value
	lock.Unlock()
	//发送全局消息
	go c.writer()
	//读取消息
	c.reader()
	defer func() {
		SuccessLogs(fmt.Sprintf("websocket客户端断开链接 %v",client))
		//清除在线的client客户端
		lock.Lock()
		delete(clientList, client)
		lock.Unlock()
		//清除绑定uid的客户端
		for k, v := range uidBindClient {
			lock.Lock()
			n := delSlice(v, client)
			lock.Unlock()
			if len(n) == 0 {
				lock.Lock()
				delete(uidBindClient, k)
				lock.Unlock()
				continue
			} else {
				lock.Lock()
				uidBindClient[k] = n
				lock.Unlock()
				continue
			}
		}
	}()
}
func (c *connection) writer() {
	for message := range c.sc {
		c.ws.WriteMessage(websocket.TextMessage, message)
	}
	c.ws.Close()
}

func (c *connection) reader() {
	for {
		_, message, err := c.ws.ReadMessage()
		if err != nil {
			h.r <- c
			break
		}
		h.b <- message
		//这里需要将客户端发来的消息进行http转发
		SuccessLogs(fmt.Sprintf("收到websocket客户端消息 %v",string(message)))
	}
}
func WebSocketRun() {
	//全局数据交互通道
	go h.Run()
	//离线消息交互
	go LogoutMasRun()

	fmt.Println("websocket start " + " ws://0.0.0.0:12358")
	http.HandleFunc("/", Handle)
	if err := http.ListenAndServe("0.0.0.0:12358", nil); err != nil {
		fmt.Println("err:", err)
	}
}

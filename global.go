package main

import (
	"encoding/json"
	"fmt"
	_ "fmt"
)

var h = hub{
	c: make(map[*connection]bool),
	u: make(chan *connection),
	b: make(chan []byte, 512),
	r: make(chan *connection),
}

type hub struct {
	c map[*connection]bool
	b chan []byte
	r chan *connection
	u chan *connection
}

var u = us{
	m: make(map[string]*connection),
}

type us struct {
	m map[string]*connection //cli => m 存储指定客户端消息管道
}

//在线的client客户端
var clientList = make(map[string]string)

//uid 和client 的绑定关系
var uidBindClient = make(map[string][]string)

func (h *hub) Run() {
	for {
		select {
		case c := <-h.r:
			h.c[c] = true
			c.Data.Ip = c.ws.RemoteAddr().String()
			c.Data.Client = c.client
			c.Data.Type = "handshake"
			//c.Data.UserList = user_list
			data, _ := json.Marshal(c.Data)
			c.sc <- data
		case c := <-h.u:
			if _, ok := h.c[c]; ok {
				delete(h.c, c)
				close(c.sc)
			}
		case data := <-h.b:
			for c := range h.c {
				fmt.Println("c", c)
				select {
				case c.sc <- data:
				default:
					delete(h.c, c)
					close(c.sc)
				}
			}
		}
	}
}

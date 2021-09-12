package main

import (
	"encoding/json"
	_ "fmt"
	"sync"
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

//存储uid的离线消息
var uidLogoutMsg = make(map[string] []string)

//引入锁
var lock sync.Mutex

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

func LogoutMasRun()  {
	for  {
		for k,v := range uidLogoutMsg  {
			lock.Lock()
			client,ok:= uidBindClient[k]
			lock.Unlock()
			if !ok {
				//fmt.Println("没有指定客户端！！")
				continue
			}else {
				//fmt.Println("有指定客户端！！")
				for _,vs := range client {
					//为了并发安全加锁
					lock.Lock()
					c ,oks := u.m[vs]
					lock.Unlock()
					if !oks {
						//fmt.Println("没有户端上线！！")
						continue
					}else {
						if len(v) > 0 {
							for _,vm := range v{
								//fmt.Println("发送消息给uid上线的客户端！！")
								c.sc <- []byte(vm)
								continue
							}
						}
					}
				}
				lock.Lock()
				delete(uidLogoutMsg,k)
				lock.Unlock()
			}

		}
	}
}

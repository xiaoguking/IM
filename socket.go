package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
)

type Msg struct {
	Cmd    int    `json:"cmd"`
	Body   string `json:"body"`
	Client string `json:"client"`
	Uid    string `json:"uid"`
}

func process(conn net.Conn) {
	defer conn.Close() // 关闭连接
	msg := &Msg{}
	for {
		reader := bufio.NewReader(conn)
		var buf [512]byte
		n, err := reader.Read(buf[:]) // 读取数据
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client端发来的数据：", recvStr)
		json.Unmarshal(buf[:n], &msg)
		m, _ := json.Marshal(msg)
		switch {
		case msg.Cmd == 1: //发送全局广播消息
			h.b <- m
			w, _ := conn.Write([]byte("全局广播发送成功")) // 发送数据
			fmt.Println(w)
		case msg.Cmd == 2: //给指定客户端发消息
			if _, ok := u.m[msg.Client]; ok { //判断客户端是否在线
				client := u.m[msg.Client]
				client.sc <- m
				w, _ := conn.Write([]byte("指定客户端发送成功")) // 发送数据
				fmt.Println(w)
			} else {
				w, _ := conn.Write([]byte("指定客户端不在线")) // 发送数据
				fmt.Println(w)
			}
		case msg.Cmd == 3: //获取在线的client客户端
			clientList, _ := json.Marshal(clientList)
			w, _ := conn.Write([]byte(string(clientList))) // 发送数据
			fmt.Println(w)
		case msg.Cmd == 4: //将uid绑定到指定客户端上
			clientId := msg.Client
			Uid := msg.Uid
			if clientId != "" && Uid != "" {
				key := Uid
				value, ok := uidBindClient[key]
				if !ok {
					value = make([]string, 0, 1)
				}
				value = append(value, clientId)
				uidBindClient[key] = value
				w, _ := conn.Write([]byte("绑定成功")) // 发送数据
				fmt.Println(w)
			}
		case msg.Cmd == 5: //向指定uid发消息
			uid := msg.Uid
			if _, ok := uidBindClient[uid]; !ok {
				w, _ := conn.Write([]byte("uid 没有在线的客户端")) // 发送数据
				fmt.Println(w)
				break
			}
			client_list := uidBindClient[uid] //uid 绑定的client
			fmt.Println("uid 绑定客户端:", client_list)
			for _, v := range client_list {
				fmt.Println("v", v)
				client, ok := clientList[v]
				if !ok {
					w, _ := conn.Write([]byte("uid 没有在线的客户端")) // 发送数据
					fmt.Println(w)
					continue
				}
				if _, ok := u.m[client]; ok { //判断客户端是否在线
					client := u.m[v]
					client.sc <- m
					w, _ := conn.Write([]byte("指定客户端发送成功")) // 发送数据
					fmt.Println(w)
				}
				w, _ := conn.Write([]byte("uid 没有在线的客户端")) // 发送数据
				fmt.Println(w)
			}
		case msg.Cmd == 6:
			list, _ := json.Marshal(uidBindClient)
			w, _ := conn.Write([]byte(list)) // 发送数据
			fmt.Println(w)
		default:
			w, _ := conn.Write([]byte("消息类型错误")) // 发送数据
			fmt.Println(w)
		}

	}
}
func SocketRun() {
	listen, err := net.Listen("tcp", "0.0.0.0:12356")

	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	fmt.Println("socket start ============" + " tcp://0.0.0.0:12356")
	for {
		conn, err := listen.Accept() // 建立连接
		if err != nil {
			fmt.Println("连接异常")
			continue
		}
		go process(conn) // 启动一个goroutine处理连接
	}
}

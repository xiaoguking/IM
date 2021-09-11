package main

func main() {
	//全局数据交互通道
	go h.Run()
	//离线消息交互
	go LogoutMasRun()
	//启动socket
	go SocketRun()
	//启动webSocket
	WebSocketRun()
}

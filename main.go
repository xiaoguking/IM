package main

func main() {
	//启动socket
	go SocketRun()
	//启动webSocket
	WebSocketRun()

}

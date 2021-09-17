package message

import (
	"encoding/json"
	"fmt"
	"net"
	"time"
)

// 向单个用户发送数据
const CMD_CLIENT_SEND_TO_ONE = 1

// 向所有用户发送数据
const CMD_SEND_TO_ALL = 2

// 获取在线状态
const CMD_GET_ALL_CLIENT = 3

// client_id绑定到uid
const CMD_BIND_UID = 4

// 解绑
const CMD_UNBIND_UID = 5

// 向uid发送数据
const CMD_SEND_TO_UID = 6

// 根据uid获取在想的clientid
const CMD_GET_CLIENT_ID_BY_UID = 7

// 判断是否在线
const CMD_IS_ONLINE = 8

// 发踢出用户
// 1、如果有待发消息，将在发送完后立即销毁用户连接
// 2、如果无待发消息，将立即销毁用户连接
const CMD_KICK = 9

// 发送立即销毁用户连接
const CMD_DESTROY = 10

// 加入组
const CMD_JOIN_GROUP = 11

// 离开组
const CMD_LEAVE_GROUP = 12

// 向组成员发消息
const CMD_SEND_TO_GROUP = 13

// 获取组成员
const CMD_GET_CLIENT_SESSIONS_BY_GROUP = 14

// 获取组在线连接数
const CMD_GET_CLIENT_COUNT_BY_GROUP = 15

// 获取在线的群组ID
const CMD_GET_GROUP_ID_LIST = 16

// 取消分组
const CMD_UNGROUP = 17

// 心跳
const CMD_PING = 201

type Msg struct {
	Cmd    int    `json:"cmd"`    //消息指令
	Body   Body   `json:"body"`   //消息实体
	Client string `json:"client"` //指定的客户端标识码
	Uid    string `json:"uid"`    //指定的uid
	Group  string `json:"group"`  //群组id
}

type Body struct {
	Type      int    `json:"type"`      //消息类型
	User      string `json:"user"`      //发送者的uid
	Content   string `json:"content"`   //消息文本内容消息
	Time      string `json:"time"`      //发送时间
	Extension string `json:"extension"` //扩展数据
	Image     string `json:"image"`     //图片消息
}

type Ret struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

var (
	m Msg
	b Body
)

// 发送广播消息
// string content 消息内容
func ToAll(content string, types int, ex string) bool {
	m.Cmd = CMD_SEND_TO_ALL
	b.Content = content
	b.Extension = ex
	b.Time = time.Now().Format("2006-01-02 15:04:05")
	b.Type = types
	b.User = "admin"
	m.Body = b
	r := send(m)
	if r.Code != 0 {
		return false
	}
	return true
}

//获取所有在线的client
// return bool
func GetClientAll() map[string]string {
	m.Cmd = CMD_GET_ALL_CLIENT
	b.Time = time.Now().Format("2006-01-02 15:04:05")
	m.Body = b
	d := send(m)
	if d.Code != 0 {
		return map[string]string{}
	}
	return d.Data.(map[string]string)
}
func send(msg Msg) Ret {
	var data Ret
	conn, err := net.Dial("tcp", "127.0.0.1:12356")
	if err != nil {
		return Ret{Code: 1, Msg: fmt.Sprintf("连接异常error %v", err), Data: ""}
	}
	d, _ := json.Marshal(msg)
	fmt.Println(string(d))
	_, err = conn.Write(d)
	if err != nil {

		return Ret{Code: 1, Msg: fmt.Sprintf("发送失败 error：%v", err), Data: ""}
	}
	buf := [512]byte{}
	n, err := conn.Read(buf[:])
	err = json.Unmarshal(buf[:n], &data)
	if err != nil {
		return Ret{Code: 1, Msg: fmt.Sprintf("发送失败 error：%v", err), Data: ""}
	}
	fmt.Println(string(buf[:n]))
	return data
}

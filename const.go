package main

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
const CMD_IS_ONLINE = 11

// 发踢出用户
// 1、如果有待发消息，将在发送完后立即销毁用户连接
// 2、如果无待发消息，将立即销毁用户连接
const CMD_KICK = 7

// 发送立即销毁用户连接
const CMD_DESTROY = 8

// 加入组
const CMD_JOIN_GROUP = 9

// 离开组
const CMD_LEAVE_GROUP = 10

// 向组成员发消息
const CMD_SEND_TO_GROUP = 11

// 获取组成员
const CMD_GET_CLIENT_SESSIONS_BY_GROUP = 12

// 获取组在线连接数
const CMD_GET_CLIENT_COUNT_BY_GROUP = 13

// 获取在线的群组ID
const CMD_GET_GROUP_ID_LIST = 14

// 取消分组
const CMD_UNGROUP = 15

// 心跳
const CMD_PING = 201


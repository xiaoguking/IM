<?php
class message
{
    // 向单个用户发送数据
    const CMD_CLIENT_SEND_TO_ONE = 1;

    // 向所有用户发送数据
    const CMD_SEND_TO_ALL = 2;

    // 获取在线状态
    const CMD_GET_ALL_CLIENT = 3;

    // client_id绑定到uid
    const CMD_BIND_UID = 4;

    // 解绑
    const CMD_UNBIND_UID = 5;

    // 向uid发送数据
    const CMD_SEND_TO_UID = 6;

    // 根据uid获取在线的clientid
    const CMD_GET_CLIENT_ID_BY_UID = 7;

    // 判断是否在线
    const CMD_IS_ONLINE = 8;
    // 发踢出用户
    // 1、如果有待发消息，将在发送完后立即销毁用户连接
    // 2、如果无待发消息，将立即销毁用户连接
    const CMD_KICK = 9;

    // 发送立即销毁用户连接
    const CMD_DESTROY = 10;

    // 加入组
    const CMD_JOIN_GROUP = 11;

    // 离开组
    const CMD_LEAVE_GROUP = 12;

    // 向组成员发消息
    const CMD_SEND_TO_GROUP = 13;

    // 获取组成员
    const CMD_GET_CLIENT_SESSIONS_BY_GROUP = 14;

    // 获取组在线连接数
    const CMD_GET_CLIENT_COUNT_BY_GROUP = 15;

    // 获取在线的群组ID
    const CMD_GET_GROUP_ID_LIST = 16;

    // 取消分组
    const CMD_UNGROUP = 17;

    // 心跳
    const CMD_PING = 201;


    const SOCKET_HOST = "127.0.0.1:12356";


    public static function send($msg){
        //使用 stream_socket_client 打开 tcp 连接
        @$fp = stream_socket_client("tcp://".self::SOCKET_HOST,$err,$errs,3);
        //向句柄中写入数据
        @fwrite($fp, $msg);
        @$ret = fread($fp, 1024);
        //关闭句柄
        @fclose($fp);
        var_dump($ret);
    }
    //封装消息
    public static function enMsg($cmd,$body = null,$client = null,$uid = null){
        $data = [
            'cmd'     => $cmd,
            'body'    => $body,
            'client'  => $client,
            'uid'     => $uid
        ];
        return json_encode($data);
    }
    /**
     * 获取与 uid 绑定的 client_id 列表
     *
     * @param string $uid
     * @return array
     */
    public static function getClientIdByUid($uid)
    {

    }

    /**
     * 向当前客户端连接发送消息
     *
     * @param string $message
     * @return bool
     */
    public static function sendToCurrentClient($client,$message)
    {

    }

    /**
     * 向所有客户端连接发送广播消息
     *
     * @param string $message 向客户端发送的消息
     * @return void
     * @throws Exception
     */
    public static function sendToAll($message)
    {
        $body = [
                    'type' => 2,
                    'user' => '系统通知',
                    'content' => $message,
                    'time' => date("Y-m-d H:i:s",time())
                ];
        $msg = self::enMsg(self::CMD_SEND_TO_ALL,$body);
        self::send($msg);
    }

    /**
     * 将 client_id 与 uid 解除绑定
     *
     * @param int $client_id
     * @param int|string $uid
     * @return void
     */
    public static function unbindUid($client_id, $uid)
    {

    }

    /**
     * 向所有 uid 发送
     *
     * @param int|string|array $uid
     * @param string $message
     *
     * @return void
     */
    public static function sendToUid($uid, $message)
    {
        foreach ($uid as $value){
            $body = [
                'type' => 1,
                'user' => 'admin',
                'content' => $message,
                'time' => date("Y-m-d H:i:s",time())
            ];
            $msg = self::enMsg(self::CMD_SEND_TO_UID,$body,null,$value);
            self::send($msg);
        }
    }
    /**
     * 踢掉某一个客户端
     *
     * @param int|string| $client
     * @return void
     */
    public static function stopClient($client)
    {
        $msg = self::enMsg(self::CMD_KICK,[],$client);
        self::send($msg);
    }
}
$obj = new message();
// $obj::sendToUid(["123","5","6"],"测试1231232131");
// $obj::sendToAll("消息开通成功");
$obj::stopClient("9d6e62771c06d8fa328e86a88bdede24");
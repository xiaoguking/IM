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

    // 根据uid获取在想的clientid
    const CMD_GET_CLIENT_ID_BY_UID = 7;

    // 判断是否在线
    const CMD_IS_ONLINE = 11;

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
}
$obj = new message();

$data = $obj::enMsg($obj::CMD_SEND_TO_ALL);
$obj::send($data);

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var logfile *os.File

var conf *Config

func init()  {
	var logpath string
	var confpath string

	flag.StringVar(&logpath, "l", "./im.log", "日志文件路径")
	flag.StringVar(&confpath, "c", "./im.conf", "配置文件路径")
	// 解析命令行参数写入注册的flag里
	flag.Parse()

	logFile, err := os.OpenFile(logpath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	logfile = logFile

	myConfig := new(Config)
	myConfig.InitConfig(confpath)
	conf = myConfig

}
func SuccessLogs(data string)  {
	log.SetOutput(logfile)
	log.SetFlags( log.Ldate)
	log.SetPrefix("[Success]")
	log.Println(data)
}
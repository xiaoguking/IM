package main

import (
	"fmt"
	"os"
	"log"
)

var logfile *os.File

func init()  {
	logFile, err := os.OpenFile("./im.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	logfile = logFile
}
func SuccessLogs(data string)  {
	log.SetOutput(logfile)
	log.SetFlags( log.Ldate)
	log.SetPrefix("[Success]")
	log.Println(data)
}
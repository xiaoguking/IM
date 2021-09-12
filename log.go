package main

import (
	"log"
)

func SuccessLogs(data string)  {
	log.SetOutput(logfile)
	log.SetFlags( log.Ldate)
	log.SetPrefix("[Success]")
	log.Println(data)
}
func ErrorLogs(data string)  {
	log.SetOutput(logfile)
	log.SetFlags( log.Ldate)
	log.SetPrefix("[Error]")
	log.Println(data)
}
package main

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"io"
)
type Response struct {
	Code int          `json:"code"`
	Msg  string       `json:"msg"`
	Data interface{}  `json:"data"`
}

var r = Response{}

//生成32位md5字串
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

//生成Guid字串
func UniqueId() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return GetMd5String(base64.URLEncoding.EncodeToString(b))
}

//删除切片中的指定元素
func delSlice(slice []string, val string) []string {
	count := len(slice)
	if count == 0 {
		return slice
	}
	if count == 1 && slice[0] == val {
		return []string{}
	}
	if count == 1 && slice[0] != val {
		return slice
	}
	var newSlice = []string{}
	for i := range slice {
		if slice[i] == val && i == count {
			return slice[:count]
		} else if slice[i] == val {
			newSlice = append(slice[:i], slice[i+1:]...)
			break
		}
	}
	return newSlice
}
func Success(data interface{}) string {
	r.Code = 0
	r.Msg = ""
	r.Data = data
	d,_ := json.Marshal(r)
	return string(d)
}

func Error(msg string) string {
	r.Code = 1
	r.Msg = msg
	r.Data = nil
	d,_ := json.Marshal(r)
	return string(d)
}
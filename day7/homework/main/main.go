package main

import (
	"go_learning_new/day7/homework/go_ini"
)

type MysqlConnection struct {
	Address  string `init:address`
	Port     int    `init:port`
	Username string `init:username`
	Password string `init:password`
}

func main() {
	var mysqlConfig MysqlConnection
	go_ini.ParseIni("./config.ini", mysqlConfig)
}

// 0. 参数的校验
// 0.1 传进来的data 必须是指针类型
// 0.2 传进来的data 必须是结构体类型
// 1. 读取文件得到字节类型数据
// 2. 一行一行读取数据
// 2.1 注释的话跳过
// 2.2 如果是[ 开头表示是节
// 2.3 如果不是[开头 用= 进行分割

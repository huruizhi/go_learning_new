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

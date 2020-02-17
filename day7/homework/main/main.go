package main

type MysqlConnection struct {
	Address  string `init:address`
	Port     int    `init:port`
	Username string `init:username`
	Password string `init:password`
}

func loadIni(v interface{}) {

}

func main() {
	var mysqlConfig MysqlConnection
	loadIni(&mysqlConfig)

}

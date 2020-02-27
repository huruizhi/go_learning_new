package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type user struct {
	id   int
	name string
	age  int
}

func DBInit() (err error) {
	dsn := "newhope:ozN0sKv@tcp(10.60.50.9:3318)/ttt"
	db, err = sql.Open("mysql", dsn)

	db.SetMaxOpenConns(100) //设置连接池的最大连接数量
	db.SetMaxIdleConns(10)  //设置最大空闲连接数，
	if err != nil {
		fmt.Println(err)
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("连接数据库成功")
	return
}

//单行查询
func queryRow(n int) {
	selectstr := "select * from user where id=?;"
	var res user
	rowObj := db.QueryRow(selectstr, n)
	rowObj.Scan(&res.id, &res.name, &res.age) //必须对rowObj进行scan 来关闭连接，不然会一直占用连接池的数据库连接数量。
	fmt.Println(res)
}

func main() {
	DBInit()
	queryRow(1)
	queryRow(2)
}

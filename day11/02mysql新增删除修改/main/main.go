package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// 数据库连接
var db *sql.DB

type user struct {
	id int
	name string
	age int
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


func modify(sql string,a ...interface{}) {
	rest, err := db.Exec(sql, a...)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rest.RowsAffected())
	fmt.Println(rest.LastInsertId())
}

// 预处理
func prepareInsertDemo() {
	sqlStr :=`insert into user(name,age) values (?,?) `
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预处理失败",err)
	}
	defer stmt.Close()
	_, err = stmt.Exec("AAA",10)
	if err != nil {
		fmt.Println("执行失败")
	}
	_, err = stmt.Exec("BBB",11)
	if err != nil {
		fmt.Println("执行失败")
	}
	fmt.Println("Insert data success")
}

func prepareSlectDemo(n int) {
	err := DBInit()
	sqlStr := `select * from user where id > ?`
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预处理失败",err)
	}
	rows, err := stmt.Query(n)
	if err != nil {
		fmt.Println("执行查询失败！",err)
	}
	defer rows.Close()
	var u user
	for rows.Next(){
		rows.Scan(&u.id,&u.name,&u.age)
		fmt.Println(u)
	}

}

func main() {
	err := DBInit()
	if err != nil {
		return
	}
	u := user{
		name:"TTBoy",
		age:3,
	}
	sql1 := `insert into user(name,age) values (?,?) `
	// sql2 := `delete from user where name=?`
	sql3 := `update user set age=0 where name=?`
	modify(sql1, u.name,u.age)
	modify(sql3, u.name)

	prepareInsertDemo()
	prepareSlectDemo(0)

}

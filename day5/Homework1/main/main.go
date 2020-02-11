package main

import (
	"fmt"
	"go_learning_new/day5/Homework1/model"
)

const (
	list = `
1. 添加学生
2. 删除学生
3. 打印所有学生
4. 退出
`
	stu1 = `stuManege := model.NewstuMgr()
	stu1 := model.NewStudent("TTboy","111111",2)
	stu2 := model.NewStudent("YDD","111112",18)
	stuManege.AddStudent("TTboy", stu1)
	stuManege.AddStudent("YDD", stu2)
	stuManege.ShowStuent()`
)

func addStudent() (name string, id string, age int) {
	fmt.Println("请输入学生姓名:")
	fmt.Scanln(&name)
	fmt.Println("请输入学生id:")
	fmt.Scanln(&id)
	fmt.Println("请输入学生年龄:")
	fmt.Scanln(&age)
	return
}

func delStudent() (name string) {
	fmt.Println("请输入学生姓名:")
	fmt.Scanln(&name)
	return
}

func main() {
	stuManege := model.NewstuMgr()
	for {
		var num int
		fmt.Println(list)
		fmt.Println("请输入编号:")
		fmt.Scanln(&num)
		if num == 1 {
			name, id, age := addStudent()
			stu := model.NewStudent(name, id, age)
			stuManege.AddStudent(name, stu)
		} else if num == 2 {
			name := delStudent()
			stuManege.DelStuent(name)
		} else if num == 3 {
			stuManege.ShowStuent()
		} else if num == 4 {
			break
		}
	}
}

package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 10
	r := reflect.ValueOf(a)
	fmt.Println(r.Kind())

	type student struct {
		Name string `json:name`
		Age  int    `jasn:age`
	}

	stu1 := student{
		Name: "DD",
		Age:  18,
	}

	t := reflect.TypeOf(stu1)
	fmt.Println(t.Name(), t.Kind())

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("name:%v, index:%v, type%v, tag:%v \n", field.Name, field.Index, field.Type, field.Tag)
	}

}

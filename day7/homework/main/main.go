package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type MysqlCon struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
	DBname   string `ini:"dbname"`
}

type RedisCon struct {
	Address string `ini:"address"`
	Port    int    `ini:"port"`
}

type iniConfig struct {
	MysqlCon `ini:"mysql"`
	RedisCon `ini:"redis"`
}

func LoadIni(filepath string, v interface{}) (err error) {
	// 0. 参数的校验
	// 0.1 传进来的data 必须是指针类型
	// 0.2 传进来的data 必须是结构体类型
	rv := reflect.ValueOf(v)
	rt := reflect.TypeOf(v)
	fmt.Println(rt.Kind())
	fmt.Println(rv.Elem().Kind())
	if rt.Kind() != reflect.Ptr {
		err = errors.New("type error: 不是一个指针")
		return
	}

	if rt.Elem().Kind() != reflect.Struct {
		err = errors.New("type error: 不是struct 类型")
		return
	}

	// 1. 读取文件得到字节类型数据
	// 2. 一行一行读取数据
	fileObj, err := os.OpenFile(filepath, os.O_RDONLY, 0644)
	if err != nil {
		return err
	}
	defer fileObj.Close()
	scanner := bufio.NewScanner(fileObj)

	var structObj reflect.Value
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		// 注释的话跳过
		if strings.HasPrefix(line, "#") || strings.HasPrefix(line, ";") {
			continue
		}
		// 空行跳过
		if len(line) == 0 {
			continue
		}
		// 获取节点
		var iniStageName string
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			stage := strings.TrimSpace(line[1 : len(line)-1])
			iniStageName = stage
			for i := 0; i < rv.Elem().NumField(); i++ {
				field := rt.Elem().Field(i)
				if field.Tag.Get("ini") == iniStageName {
					structObj = rv.Elem().Field(i)
					continue
				}
			}
		}

		// 获取属性与值
		// 判断是否包含"="
		if strings.Contains(line, "=") {

			// 判断key value 是否为空
			key := strings.Split(line, "=")[0]
			value := strings.Split(line, "=")[1]
			if key == "" || value == "" {
				continue
			}

			rv := structObj
			rt := structObj.Type()
			// 分析结构体
			for i := 0; i < rv.NumField(); i++ {
				// 获取结构体的tag 并与key 进行比较
				if rt.Field(i).Tag.Get("ini") == key {

					//将value 赋值给 结构体的对应变量
					// 判断结构体的变量类型
					switch rv.Field(i).Kind() {
					case reflect.String:
						rv.Field(i).SetString(value)
					case reflect.Int:
						num, err := strconv.ParseInt(value, 10, 0)
						if err != nil {
							err = fmt.Errorf("字段 %s 为 int 型", rt.Field(i).Name)
							return err
						}
						rv.Field(i).SetInt(num)
					}
				}
			}
		}
	}

	return
}

func main() {
	var inicon iniConfig
	err := LoadIni("./config.ini", &inicon)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v\n", inicon.MysqlCon)
	fmt.Printf("%#v\n", inicon.RedisCon)

}

// 2.1 注释的话跳过
// 2.2 如果是[ 开头表示是节
// 2.3 如果不是[开头 用= 进行分割

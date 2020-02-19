package go_ini

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"
)

func ParseIni(filePath string, v interface{}) error {
	file := fileIsValid(filePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	rv := reflect.ValueOf(v)
	rt := reflect.TypeOf(v)
	err := errors.New("type error")
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return err
	}

	mysqlConfig := true
	for scanner.Scan() {
		line := scanner.Text() // scanner.Bytes()
		if line == "[mysql]" {
			continue
		}
		if mysqlConfig {
			k := strings.Split(line, "=")[0]
			v := strings.Split(line, "=")[1]
			for i := 0; i < rt.NumField(); i++ {
				field := *rt.Field(i)
				if field.Tag.Get("ini") == k {
					filedName := field.Name
					file := rv.Elem().FieldByName(filedName)
					fmt.Printf("%T %s", file, string(v))
				}
			}
		}

	}

	return nil
}

func fileIsValid(filePath string) (f *os.File) {
	f, err := os.OpenFile(filePath, os.O_RDONLY, 644)
	if err != nil {
		panic(err)
	}
	return f
	json.Unmarshal()
}

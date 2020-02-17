// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"path"
	"runtime"
)

func main() {
	//var x bitArry.IntSet
	//x.Add(5)
	//x.Add(6)
	//x.Add(67)
	//fmt.Println(x.Words)
	//fmt.Println(x.String())
	//a := 2
	//b := 1 << 1
	//fmt.Println(a, b)
	//if b&(1<<1) != 0 {
	//	println("OK")
	//}

	//var t time.Weekday
	//t = 0
	//fmt.Println(t.String())
	//
	//loc1, _ := time.LoadLocation("Asia/Shanghai")
	//loc2, _ := time.LoadLocation("Europe/Berlin")
	//fmt.Println(loc1)
	//fmt.Println(loc2)
	//const longForm = "Jan 2, 2006 at 3:04pm"
	//t1,_ := time.ParseInLocation(longForm,"Jul 9, 2012 at 5:02am", loc1)
	//t2,_ := time.ParseInLocation(longForm,"Jul 9, 2012 at 5:02am", loc2)
	//fmt.Println(t1)
	//fmt.Println(t2)
	//fmt.Println(t2.Location())
	//now := time.Now()
	//fmt.Println(now.Location())

	pc, file, line, ok := runtime.Caller(0)
	if !ok {
		panic("call runtime failed!")
	}
	fmt.Println(runtime.FuncForPC(pc).Name())
	fmt.Println(path.Base(file))
	fmt.Println(line)
}

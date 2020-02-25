package main

import (
	"sync"
)

var icons map[string]string

var loadOnce sync.Once

func loadIcons() {
	icons = map[string]string{
		"left":  "left.png",
		"right": "right.png",
		"up":    "up.png",
		"down":  "down.png",
	}
}

func Icon(name string) string {
	loadOnce.Do(loadIcons) // 并发时icons 初始化一次
	return icons[name]
}

func main() {
	for i := 0; i < 10; i++ {
		go Icon("up") // 并发时icons 初始化一次
	}
}

package split

import (
	"strings"
	"time"
)

func Split(str string, sep string) (splitList []string) {
	sepLenth := len(sep)
	for {
		time.Sleep(time.Millisecond * 50)
		num := strings.Index(str, sep)
		if num >= 0 {
			splitList = append(splitList, str[:num])
			str = str[num+sepLenth:]
		} else {
			splitList = append(splitList, str)
			break
		}
	}

	return
}

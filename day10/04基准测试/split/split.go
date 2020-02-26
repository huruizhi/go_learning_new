package split

import (
	"strings"
)

func Split(str string, sep string) []string {
	sepLenth := len(sep)
	splitList := make([]string, strings.Count(str, sep)+1)
	num := strings.Index(str, sep)
	for num >= 0 {
		//time.Sleep(time.Millisecond*50)
		splitList = append(splitList, str[:num])
		str = str[num+sepLenth:]
		num = strings.Index(str, sep)
	}

	splitList = append(splitList, str)
	return splitList
}

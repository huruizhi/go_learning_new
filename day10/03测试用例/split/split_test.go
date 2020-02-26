package split

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	type testCase struct {
		str  string
		sep  string
		want []string
	}
	// 测试组
	testGroup := map[string]testCase{
		"test1": {str: ",,abcd,,efg,,hijk", sep: ",,", want: []string{"", "abcd", "efg", "hijk"}},
		"test2": {str: "a,b,c,d", sep: ",", want: []string{"a", "b", "c", "d"}},
	}
	for k, v := range testGroup {
		// 调用测试用例
		// go test -run=test2 调用单独的测试用例
		t.Run(k, func(t *testing.T) {
			splitList := Split(v.str, v.sep)
			if !reflect.DeepEqual(splitList, v.want) {
				t.Errorf("want: %#v but got:%#v\n", v.want, splitList)
			}
		})
	}
}

// 测试覆盖率 go test --cover
// 输出到文件 go test --cover -coverprofile=cc.out
// 通过网页打开文件  go tool cover  -html=cc.out

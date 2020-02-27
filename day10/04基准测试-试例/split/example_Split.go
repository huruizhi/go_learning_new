package split

import "fmt"

func exampleSplit() {
	fmt.Println(Split(",,abcd,,efg,,hijk", ",,"))
	// Output:
	// [abcd efg hijk]
}

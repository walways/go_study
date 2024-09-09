package geektime

import (
	"fmt"
	"testing"
)

func TestMap1(t *testing.T) {

	var badMap2 = map[interface{}]int{
		"1":      1,
		[]int{2}: 2, // 这里会引发panic。
		3:        3,
	}

	fmt.Println(badMap2)

}

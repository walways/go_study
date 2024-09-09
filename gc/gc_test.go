package gc

import (
	"fmt"
	"runtime"
	"testing"
)

func TestGc1(t *testing.T) {
	type MyStruct struct {
		value int
	}
	// 分配在堆上的变量
	a := new(int)
	*a = 10

	// 分配在堆上的结构体
	b := &MyStruct{value: 20}

	// 分配在栈上的变量
	c := 30

	fmt.Println("a:", *a)      // 输出: a: 10
	fmt.Println("b:", b.value) // 输出: b: 20
	fmt.Println("c:", c)       // 输出: c: 30

	// 强制进行垃圾回收
	runtime.GC()

	fmt.Println("a:", *a)      // 输出: a: 10
	fmt.Println("b:", b.value) // 输出: b: 20
	fmt.Println("c:", c)       // 输出: c: 30

}

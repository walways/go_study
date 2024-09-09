package base

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestUnsafe(t *testing.T) {
	x := 43
	point := unsafe.Pointer(&x)
	sx := (*float64)(point)
	fmt.Printf("aaaa: %+v", *sx)

	test := struct {
		Name string `json:"name"`
	}{
		Name: "dasd",
	}
	point1 := unsafe.Pointer(&test)
	sx1 := *(*string)(point1)

	fmt.Printf("bbbb: %+v", sx1)
}

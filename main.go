package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	sl := make([]int, 0)

	for i := 1; i <= 5; i++ {
		sl = append(sl, i)
	}
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&sl))
	fmt.Println("mem region = ", sh.Data, "\tpointer = ", &sh.Data)

	fmt.Println("foo exec")
	foo(sl)
}

func foo(sl []int) {
	sl[4] = 10
	sb := (*reflect.SliceHeader)(unsafe.Pointer(&sl))
	fmt.Println("mem region = ", sb.Data, "\tpointer = ", &sb.Data)
}

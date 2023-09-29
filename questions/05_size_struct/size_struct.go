package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Printf("unsafe.Sizeof(struct{}{}): %v\n", unsafe.Sizeof(struct{}{}))

}

// unsafe.Sizeof(struct{}{}): 0

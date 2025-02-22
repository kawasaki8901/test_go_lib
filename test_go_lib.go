package main

import (
	"fmt"

	"github.com/kawasaki8901/test_go_lib/lib"
)

func main() {
	fmt.Printf("call my lib\n")
	lib.Hello()
}

func Call() {
	fmt.Printf("call my lib\n")
	lib.Hello()
}

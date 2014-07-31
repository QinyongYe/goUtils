package main

import "fmt"

import (
)

func main() {
	fmt.Println(string(2))                 // illegal: 1.2 cannot be represented as an int
//	string(65.0)             // illegal: 65.0 is not an integer constant
}

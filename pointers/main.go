package main

import "fmt"

func main() {
	var x int = 1
	p := &x
	fmt.Println("address of x", p)
	*p = 2
	fmt.Println("value of x", x)

}

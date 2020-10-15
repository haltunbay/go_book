package main

import "fmt"

func main() {
	var x uint8 = 255
	fmt.Printf("%08b \n", x)
	x++
	fmt.Printf("%08b \n", x)
	x = x << 1
	fmt.Printf("%08b \n", x)
	x = 2<<6 | 1<<1
	fmt.Printf("%08b \n", x)
	x = ^x
	fmt.Printf("%08b \n", x)
	x = x >> 1
	fmt.Printf("%08b \n", x)

	f := 1e11
	fmt.Printf("%g", f)
}

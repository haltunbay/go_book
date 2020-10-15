package main

import "fmt"

var pc [256]byte

func main() {
	init1()
}

func init1() {
	for i := range pc {
		old := pc[i]
		pc[i] = pc[i/2] + byte(i&1)
		fmt.Println(i, old, pc[i], i&1)
	}
}

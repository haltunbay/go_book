package main

import (
	"fmt"
	"strings"
)

func main() {
	// x, _ := strconv.Atoi(os.Args[0])
	// y, _ := strconv.Atoi(os.Args[1])
	// fmt.Println(gcd(x, y))
	//printChars()
	testBasename()
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func printChars() {
	s := "Hello açşğüİ"
	for _, c := range s {
		fmt.Printf("%s %.12b % x %d\n", string(c), c, c, c)
	}
}

func testBasename() {
	fmt.Println(basename("a.b"))
	fmt.Println(basename("/home/haltunbay/a.b"))
	fmt.Println(basename("http:///home.haltunbay/a.b"))
}

// basename removes directory components and a .suffix.
// e.g., a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
func basename(s string) string {
	slash := strings.LastIndex(s, "/")

	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

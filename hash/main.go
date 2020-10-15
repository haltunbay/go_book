package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

var pc [256]byte

var s = flag.String("s", "SHA256", "Hash Type")
var i = flag.String("i", "", "input")

func main() {
	flag.Parse()

	switch *s {
	case "SHA256":
		fmt.Printf("'%s':\t %x\n", *i, sha256.Sum256([]byte(*i)))

	case "SHA384":
		fmt.Printf("'%s':\t%x\n", *i, sha512.Sum384([]byte(*i)))

	case "SHA512":
		fmt.Printf("'%s':\t%x\n", *i, sha512.Sum512([]byte(*i)))
	default:
		fmt.Errorf("Invalid hash type %s", *i)
	}
}

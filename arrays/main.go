package main

import "fmt"

func main() {
	months := [...]string{1: "January", 12: "December"}
	fmt.Println(months[12])
	dec := months[12:]
	fmt.Println(dec)
	dec[0] = "None"
	fmt.Println(months[12])

	s := "hello ÜĞŞÇİ"
	fmt.Println(len(s))
	fmt.Println(len([]rune((s))))

	a := [...]int{1, 2, 3}

	print(&a)
	reverse(&a)
	print(&a)

	b := []int{1, 2, 3, 4}
	rotate(b)
	rotate(b)
	rotate(b)
	rotate(b)
}

func print(a *[3]int) {
	fmt.Println(a)

}

func reverse(a *[3]int) {

	for i, j := 0, len(a)-1; i < j; {
		a[i], a[j] = a[j], a[i]
		i++
		j--
	}
}

func rotate(a []int) {
	last := a[0]

	for i := 0; i < len(a)-1; i++ {
		a[i] = a[i+1]
	}
	a[len(a)-1] = last
	fmt.Println(a)
}

func removeDuplicates(s []string) {
	out := []string{}
	for i, s1 := range s {
		if s1 == s[i+1] {
			out = append(out, s[:i], s[i:])
		}
	}
}

package leetcode

import "fmt"

func b() {
	defer un(trace("b"))
	fmt.Println("b codes")
	a()
}

func a() {
	defer un(trace("a"))
	fmt.Println("a codes")
}

func trace(s string) string {
	fmt.Println("start:", s)
	return s
}

func un(s string) string {
	fmt.Println("end:", s)
	return s
}

func Defer() {
	b()
}

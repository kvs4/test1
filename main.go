package main

import "fmt"

func sum(a, b, c int) {
	x := a + b + c
}

func sub(a, b, c int) {
	res := a - b - c
}
func main() {
	a := 1
	b := 3
	c := a + b
	fmt.Println("c =", c)
}

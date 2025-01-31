package main

import "fmt"

type A struct {
	a float64
}

type B struct {
	A
	b float64
}

func t1(input any) {
	b, ok := input.(A)
	fmt.Println(ok, b)
}

func main() {
	b := B{
		A: A{
			a: 1.0,
		},
		b: 2.0,
	}

	fmt.Println(b.a)
	t1(b)
}

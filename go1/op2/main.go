package main

import "fmt"

type A []string

func (a *A) String() string {
	ret := ""
	for _, v := range *a {
		ret += v
	}
	return ret
}

func main() {
	a := A{"a", "b", "c"}
	ret, _ := fmt.Println(a)
	fmt.Println(ret)
}

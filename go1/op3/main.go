package main

import (
	"fmt"
)

func checkIfSlice(input any) {
	if _, ok := input.([]any); ok {
		fmt.Println("input is slice")
		return
	}

	fmt.Println("not slice")
}

func main() {
	a := []any{1, 2, 3}

	checkIfSlice(a)
}

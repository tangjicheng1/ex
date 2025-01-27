package main

func t1(s ...string) {
	for _, v := range s {
		println(v)
	}
}

func t2(s ...int) {
	println(s[0])
}

func main() {
	t1()
	t1("a", "b", "c")

	t2(1, 2, 3)
	// t2()
}

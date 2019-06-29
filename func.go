package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func say() {
	println("Hello world")
}

func main() {
	k := func(x, y int) int {
		return x + y
	}
	say()
	fmt.Println(k(42, 12))
}

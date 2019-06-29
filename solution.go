package main

import "fmt"

func twofer(x string) {
	fmt.Printf("One for %v,one for me.", x)
}

func main() {
	twofer("Alice")
	// twofer("")
}

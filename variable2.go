package main

import "fmt"

var i, j int = 1, 2

func main() {
	// var c, python, java = true, false, "no!"
	// k := 3
	// fmt.Println(i, j, k, c, python, java)
	var i int
	var f float64
	var b bool
	var s string
	var r rune
	fmt.Printf("%v %v %v %v %q\n", i, f, b, r, s)
}

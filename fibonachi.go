package main

import "fmt"

func fibonacci() func() int {
	var a, b int = 0, 1
	return func() int {
		var reuslt = a + b
		a = b
		b = reuslt
		return reuslt
	}

}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		if i == 0 {
			fmt.Println(0)
		} else {
			fmt.Println(f())
		}
	}
}

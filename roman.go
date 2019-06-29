package main

import (
	"fmt"
	"math"
)

fun

func roman(num int) string {
	total := num
	str := ""
	if math.Mod(total, 100) == 0 {
		str += "C"
		total -= 100
	}
	multiple := total/10
	if multiple == 9 {
		str += "XC"
		total -= 90
	}else if(multiple ==)
	if math.Mod(total, 100) == 0 {
		str += "C"
		total -= 100
	}
	if math.Mod(total, 100) == 0 {
		str += "C"
		total -= 100
	}
	if math.Mod(total, 100) == 0 {
		str += "C"
		total -= 100
	}
	if math.Mod(total, 100) == 0 {
		str += "C"
		total -= 100
	}

}

func main() {
	for i := 0; i < 100; i++ {
		fmt.Println(roman(i))
	}
}

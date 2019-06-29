package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	arrayString := strings.Split(strings.Replace(s, ",", "", -1), " ")
	for _, x := range arrayString {
		m[x]++
	}
	return m
}

func main() {
	s := "If it looks like a duck, swims like a duck, and quacks like a duck, then it probably is a duck"
	fmt.Println(WordCount(s))
}

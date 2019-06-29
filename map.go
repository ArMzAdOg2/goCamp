package main

import "fmt"

func main() {
	var m map[string]Student
	m = make(map[string]Student)
	student := Student{"Active", "Manasak", "P6", "12", "11"}
	m["191"] = student
	elem, ok := m["qq"]
	fmt.Println(m["191"])
	fmt.Println(elem, ok)
}

type Student struct {
	StudentInformation string
	Name               string
	Class              string
	Age                string
	No                 string
}

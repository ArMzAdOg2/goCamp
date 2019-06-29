package main

import "fmt"

type Person struct {
	Name string
}

func (person Person) Walk() {
	fmt.Println(person.Name, "walking")
}

func (person Person) Eat() {
	fmt.Println(person.Name, "eating")
}

func (person Person) Geeting() {
	fmt.Println("hello", person.Name)
}

func (person Person) get() string {
	return person.Name
}

func (person *Person) set(value string) {
	person.Name = value
}

func main() {
	p := Person{"Arm"}
	p.Walk()
	p.Geeting()
	p.Eat()
	fmt.Println(p.get())
	p.set("Test")
	fmt.Println(p.get())
}

package main

import "fmt"

type Person struct {
	firstName string
	lastName string
}

type SecretAgent struct {
	Person
	licenseToKill bool
}

type human interface {
	speak()
}

func (p Person) speak() {
	fmt.Printf("Good morning, I'm %s %s\n", p.firstName, p.lastName)
}

func (sa SecretAgent) speak() {
	fmt.Printf("Good morning, my name is %s %s, I have licence to kill\n", sa.firstName, sa.lastName)
}

func saySomething(h human) {
	h.speak()
}

func main () {
	person1 := Person{
		"Elias",
		"Moon",
	}

	secretAgent1 := SecretAgent{
		Person {
			"James",
			"Bond",
		},
		true,
	}

	saySomething(person1)
	saySomething(secretAgent1)
	saySomething(secretAgent1.Person)
}
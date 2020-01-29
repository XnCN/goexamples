package main

import "fmt"

type Person struct {
	fullName string
	city     string
	age      int
}

//mapler c#daki hastable gibi anahtar değer şeklinde
func main() {
	myMap := make(map[int]Person)
	myMap[1] = Person{
		fullName: "Aydemir Eray Han",
		city:     "Kocaeli",
		age:      22,
	}
	myMap[2] =Person{
		fullName: "Irmak Han",
		city:     "Kocaeli",
		age:      12,
	}
	for index,p := range myMap{
		fmt.Println(index,"|",p.fullName,p.city,p.age)
	}
}

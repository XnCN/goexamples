package main

import "fmt"

type Person struct {
	name string
	age  int
}

type DataTransformModel struct {
	data    []Person
	channel chan []Person
}

func main() {
	persons := []Person{
		{
			name: "Aydemir Eray Han",
			age:  22,
		},
		{
			name: "Irmak Han",
			age:  12,
		},
	}
	channel:= make(chan []Person)
	go changePersonsAges(DataTransformModel{
		data:    persons,
		channel: channel,
	})
	persons= <- channel
	fmt.Println(persons)
}

func changePersonsAges(transform DataTransformModel) {
	for i := range transform.data {
		transform.data[i].age +=10
	}
	transform.channel <- transform.data
}

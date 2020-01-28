package main

import "fmt"

type City struct {
	code int
	name string
}
type Person struct {
	fullName string
	age      int
}

func main() {
	list := make(map[City]Person)
	list[City{
		code: 41,
		name: "Kocaeli",
	}] = Person{
		fullName: "Aydemir Eray Han",
		age:      22,
	}
	list[City{
		code: 34,
		name: "İstanbul",
	}] = Person{
		fullName: "Irmak Han",
		age:      12,
	}

	for city,person := range list{
		if city.code == 34{
			delete(list,city)
			return
		}
		fmt.Println(city.code,city.name,"|",person.fullName,person.age)
	}


}

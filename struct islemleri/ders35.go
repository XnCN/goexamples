package main

import "fmt"

func main(){
	hum := Human{
		age: 22,
	}
	hum.changeAge()
	hum.changeAgeWithPointer()
	fmt.Println(hum)
}

type Human struct {
	age int
}

func (h Human) changeAge() int{
	return h.age*2
}

func (h *Human) changeAgeWithPointer(){
	h.age = 1
}
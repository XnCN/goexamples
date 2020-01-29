package main

import (
	"encoding/json"
	"fmt"
)

type Customer struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	jsonString := `{"firstname":"Aydemir Eray","lastname":"Han","age":22}`
	var myCustomer Customer
	json.Unmarshal([]byte(jsonString), &myCustomer)
	fmt.Printf("Customer Name:%v,Surname:%v,Age:%v \n", myCustomer.FirstName, myCustomer.LastName, myCustomer.Age)

	//objeden json'a
	j, _ := json.Marshal(myCustomer)
	fmt.Println(string(j))
}

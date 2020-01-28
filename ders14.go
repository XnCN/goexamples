package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		mystring = "1"
		x        = 10
		f        = 2.0
	)
	fmt.Println(mystring, x, f)
	//str to int
	number, _ := strconv.Atoi(mystring)
	fmt.Println(number + x)
	//int to str
	stringnumber := strconv.Itoa(number)
	fmt.Println(stringnumber)
}

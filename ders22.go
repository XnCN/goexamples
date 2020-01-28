package main

import "fmt"

func main(){
	//slice üzerinde değişiklik yaparsak ana dizi değişir
	myArray :=[...]int{45,23,43}
	mySlice := myArray[:]
	fmt.Println(mySlice)
	mySlice[0]=11
	fmt.Println(mySlice)
	fmt.Println(myArray)
}

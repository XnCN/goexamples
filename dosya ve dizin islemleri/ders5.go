package main

import (
	"fmt"
	"log"
	"os"
)

//dosya açmak ve kapatmak

var(
	file *os.File
	err error
)
func main(){
	file,err = os.Open("dosya ve dizin islemleri/demo.txt")
	defer file.Close()
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println(file.Name())
}

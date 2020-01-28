package main

import (
	"fmt"
	"io/ioutil"
)

func main(){
	text,err :=ioutil.ReadFile("text.txt")
	if err !=nil{
		fmt.Println(err)
	}
	fmt.Println(string(text))
}

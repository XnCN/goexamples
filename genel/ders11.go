package main

import (
	"fmt"
	"io/ioutil"
)
func main(){
	path:="/Users/eray/Desktop/go/text.txt"
	text,err := ioutil.ReadFile(path)
	if err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Println(string(text))
}
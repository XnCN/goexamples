package main

import "fmt"

func main(){
	colors := []string{"orange","red","blue"}
	colors = append(colors,"black")
	fmt.Println(colors)
	colors = append(colors[1:len(colors)])
	fmt.Println(colors)
}
package main

import (
	"fmt"
	"time"
)

func main(){
	c:=make(chan byte)
	go alphabet(c)
	go printChar(c)
	time.Sleep(100*time.Millisecond)
}

func alphabet(channel chan byte){
	for i:=byte('a');i<=byte('z') ; i++  {
		channel <- i
	}
}

func printChar(channel chan byte){
	for{
		fmt.Println(string(<-channel))
	}
}
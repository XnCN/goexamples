package main

import (
	"fmt"
	"runtime"
	"time"
)

//paralel programlama
func main(){
	runtime.GOMAXPROCS(2) //kullanılacak işlemci sayısı
	go printAlphabet()
	time.Sleep(100*time.Millisecond)
}

func printAlphabet(){
	for i:=byte('a');i<=byte('z') ;i++  {
		go fmt.Println(string(i))
	}
}
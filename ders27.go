package main

import (
	"fmt"
	"time"
)
//eş zamanlı
func main() {
	go printalphabet()
	time.Sleep(100*time.Millisecond)
}

func printalphabet() {
	for i := byte('a'); i <= byte('z'); i++ {
		fmt.Println(string(i))
	}
}

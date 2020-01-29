package main

import (
	"fmt"
	"math/rand"
	"time"
)

var source = rand.NewSource(time.Now().UnixNano())
const _charset = "aqwertyuopilkjhgfdsazxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890"
func main(){
	fmt.Println(generateRandomPassword(12))
}

func generateRandomPassword(lenght int) string{
	x:= make([]byte,lenght)
	for i:=range x{
		x[i]= _charset[source.Int63() % int64(len(_charset))]
	}
	return string(x)
}
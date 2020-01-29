package main

//string verinin birleştirme işlemi performanslı olarak
//buffer'ı çok fazla veriyi birleştirmek için kullan
import (
	"bytes"
	"fmt"
)
func main(){
	var x bytes.Buffer
	for i:=0;i<100;i++{
		x.WriteString("-deneme mesajı")
	}
	fmt.Println(x.String())
}
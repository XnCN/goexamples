package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	reader:=bufio.NewReader(os.Stdin)
	fmt.Println("Adınızı girin")
	text,_ := reader.ReadString('\n')
	fmt.Println("Hoşgeldin",text)

}
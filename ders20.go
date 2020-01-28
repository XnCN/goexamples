package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Lütfen sayıyı gir:")
	number,_ := reader.ReadString('\n')
	f,err := strconv.ParseFloat(strings.TrimSpace(number),64)
	if err !=nil{
		println(err)
	}else{
		println(f)
	}

}

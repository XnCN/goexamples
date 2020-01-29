package main

import (
	"fmt"
	"log"
	"os"
)

//permission kontrol

func main(){
	file,err:=os.OpenFile("dosya ve dizin islemleri/demo.txt",os.O_RDWR,666)
	defer file.Close()
	if err !=nil{
		if os.IsPermission(err){
			log.Fatal(err)
		}
	}
	fmt.Println(file.Name())
}

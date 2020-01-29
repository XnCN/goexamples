package main

import (
	"log"
	"os"
)

var(
	file *os.File
	err error
)

func main(){
	//dosya adı , ne için açtın, okuma yazma izinleri ister 666 77 gibi ister os mode ile
	file,err = os.OpenFile("dosya ve dizin islemleri/demo.txt",os.O_APPEND,os.ModeAppend)
	defer file.Close()
	if err !=nil{
		log.Fatal(err)
	}
	file.WriteString("burası string")
}

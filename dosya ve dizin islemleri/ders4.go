package main
//dosyanın varlığını kontrol etmek

import (
	"fmt"
	"log"
	"os"
)
func main(){
	fileInfo,err := os.Stat("dosya ve dizin islemleri/demo.txt")
	if err != nil{
		if os.IsNotExist(err){
			log.Fatal("file does not exits")
		}
	}
	fmt.Println(fileInfo.Name())
}

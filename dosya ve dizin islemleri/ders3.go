package main

//dosyanın adını ve yolunu değiştirmek
import (
	"log"
	"os"
)
func main(){
	originalPath:="dosya ve dizin islemleri/demo.txt"
	newPath:="dosya ve dizin islemleri/move/demo.txt"
	err := os.Rename(originalPath,newPath)
	if err !=nil{
		log.Fatal(err)
	}
}

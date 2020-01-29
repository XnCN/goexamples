package main
//dosya oluşturma
import (
	"log"
	"os"
)
var(
	newFile *os.File
	err error
)
func main(){
	newFile,err = os.Create("dosya ve dizin islemleri/demo.txt")
	if err!=nil{
		log.Fatal(err)
	}
}
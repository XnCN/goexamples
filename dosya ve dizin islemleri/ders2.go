package main
//dosya bilgisini almak
import (
	"fmt"
	"log"
	"os"
)

var(
	fileInfo *os.FileInfo
)
func main(){
	fileInfo,err := os.Stat("dosya ve dizin islemleri/demo.txt")
	if err !=nil{
		log.Fatal(err)
	}
	fmt.Println(fileInfo.Name(),fileInfo.Size(),fileInfo.IsDir(),fileInfo.Mode())
}

package main

import (
	"fmt"
	"log"
	"os"
)

func main(){
	file,err := os.OpenFile("dosya ve dizin islemleri/demo.txt",os.O_WRONLY,666)
	defer file.Close()
	if err !=nil{
		log.Fatal(err)
	}
	stringBytes := []byte("dosyaya yazılıcak metin\n")
	writenBytes,err := file.Write(stringBytes)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Printf("%v byte yazıldı",writenBytes)

}

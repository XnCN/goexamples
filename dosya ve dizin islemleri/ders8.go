package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

//dosya kopyalama

func main(){
	originalFile,err := os.Open("dosya ve dizin islemleri/demo.txt")
	defer originalFile.Close()
	if err!=nil{
		log.Fatal(err)
	}
	//yeni dosya olustur
	newFile,err:= os.Create("dosya ve dizin islemleri/democopy.txt")
	defer newFile.Close()
	if err!=nil{
		log.Fatal(err)
	}
	//orjinal dosyayı kopyalama kopyalama hafızada yapıyor
	bytes,err :=io.Copy(originalFile,newFile)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Printf("file copyed , %d",bytes)

	//dosya içeriğini işle ve belleği diske dök
	err = newFile.Sync()
	if err!=nil{
		log.Fatal(err)
	}
}

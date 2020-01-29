package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main(){
	//gecici klasör olustur
	tempDir,err := ioutil.TempDir("","geciciders11")
	if err !=nil{
		log.Fatal(err)
	}
	fmt.Printf("Geçici klasör olusturuldu : %v \n",tempDir)
	//gecici dosya olustur
	tempFile,err := ioutil.TempFile(tempDir,"gecici.txt")
	defer tempFile.Close()
	if err !=nil{
		log.Fatal(err)
	}
	fmt.Println(tempFile.Name())

	//gecici dosya üstünde işlemler yap ve sil
	err = os.Remove(tempFile.Name())
	if err !=nil{
		log.Fatal(err)
	}
	err =os.Remove(tempDir)
	if err !=nil{
		log.Fatal(err)
	}
}
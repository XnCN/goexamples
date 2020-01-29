package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

//konsoldan girilen metni demo.txt ye ekle

func main(){
	fmt.Print("Lütfen bir metin giriniz:")
	reader := bufio.NewReader(os.Stdin)
	userInput,err := reader.ReadString('\n')
	if err!=nil{
		log.Fatal(err)
	}
	file,err := os.OpenFile("dosya ve dizin islemleri/demo.txt",os.O_WRONLY,777)
	defer file.Close()
	if err !=nil{
		log.Fatal(err)
	}
	userInputBytes := []byte(userInput)
	writenBytes,err:=file.Write(userInputBytes)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Printf("%v byte dosyaya yazıldı",writenBytes)


}

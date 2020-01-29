package main

import "fmt"

//kanallar gorutinler arasındaki iletişimi sağlar
// ch <- v  = v değişkenini ch kanalına gönder
// b:= <-ch  = b değişkenine ch kanalından gelen veriyi aktar
func main(){
	numbers := []int{1,2,3,5,12}
	channel := make(chan int)
	go sumNumbers(numbers,channel)
	fmt.Println("sayıların toplamı :",<-channel)
}

func sumNumbers(numbers []int,channel chan int){
	sum:=0
	for _,number := range numbers{ sum+=number }
	channel <- sum
}

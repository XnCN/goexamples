package main

import "fmt"

type DataTransfer struct {
	data    []int
	channel chan int
}

func main() {
	numbers := []int{1, 2, 3, 5, 12}
	c := make(chan int)
	go sumNumbersWithStruct(DataTransfer{
		data:    numbers,
		channel: c,
	})
	sum := <-c
	fmt.Println("sayıların toplamı:", sum)
}

func sumNumbersWithStruct(transfer DataTransfer) {
	sum := 0
	for _, number := range transfer.data {
		sum += number
	}
	transfer.channel <- sum
}

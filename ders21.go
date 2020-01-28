package main

import "fmt"

func main() {

	dizi := [3]int{}
	dizi[0] = 32
	dizi[1] = 23
	dizi[2] = 54
	fmt.Println(dizi)

	//renkler
	var colors [3]string
	colors[0] = "red"
	colors[1] = "green"
	colors[2] = "blue"
	fmt.Println(colors)

	numbers := [5]int{4,5,6,7,12}
	fmt.Println(numbers)
	fmt.Println("Number of numbers",len(numbers))

	//dizi sınırını ortadan kaldırmak için [...]tip

	autoindexarray := [...]int{1,2,3,4,5,6,7,8,9}
	fmt.Println(autoindexarray)
	sum:=0
	//autoindexarray toplamı
	for _,n := range autoindexarray{
		sum+=n
	}
	fmt.Println(sum)
}

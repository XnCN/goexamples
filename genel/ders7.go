package main

import "fmt"

func main() {
	dizi := [5]int{1, 2, 3, 5, 6}

	var toplam int
	for i := 0; i <= len(dizi)-1; i++ {
		toplam += dizi[i]
	}

	fmt.Println("Dizideki elemanların toplamı ", toplam)
	fmt.Println("Dizideki elemanların ortalaması ", toplam/len(dizi))
}

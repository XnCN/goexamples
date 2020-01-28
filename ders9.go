package main

import "fmt"

func main() {
	toplam := func(x, y int) int {
		return x + y
	}
	karelerToplamı := func(sayilar ...int) int {
		toplam := 0
		for _, n := range sayilar {
			toplam += n * n
		}
		return toplam
	}
	fmt.Println(toplam(1, 2))
	fmt.Println(karelerToplamı(1, 2, 3, 4, 5))
}

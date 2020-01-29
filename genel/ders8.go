package main

import "fmt"

func main() {
	toplam := sayitoplamlari(1, 2, 3, 4, 5)
	fmt.Println(toplam)
}

//variadic function
func sayitoplamlari(sayilar ...int) int {
	toplam := 0
	for _, n := range sayilar {
		toplam += n
	}
	return toplam
}

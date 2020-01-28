package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(fmt.Sprintf("%v sayısı 2 ile tam bölünüyor", i))
		}
	}
}

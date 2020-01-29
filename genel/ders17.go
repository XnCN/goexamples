package main

import "fmt"

func main() {
	message, result := add(5, 4, "Sonuç")
	fmt.Println(message, result)
}

func add(x, y int, message string) (string, int) {
	return message, x + y
}

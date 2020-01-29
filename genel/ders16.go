package main

import "fmt"

//pointer uygulaması
func main() {
	message := "Hello Eray"
	sayHello(&message)
	fmt.Println(message)
}

func sayHello(message *string) {
	fmt.Println(*message)
	*message = "Hello Go!"
}

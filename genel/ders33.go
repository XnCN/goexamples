package main

//ortam değişkenlerine erişim

import (
	"fmt"
	"os"
	"strings"
)

func main(){
	envorimets := make(map[string]string)
	for _,env := range os.Environ(){
		data := strings.Split(env,"=")
		envorimets[data[0]] = data[1]
	}
	fmt.Println(envorimets["USER"])

	//kısa yol
	fmt.Println(os.Getenv("USER"))
	fmt.Println(os.Getenv("GOROOT"))
}
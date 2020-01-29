package main

import (
	"fmt"
	"time"
)

func main() {
	switch time.Now().Weekday() {
	case time.Monday:
		fmt.Println("Pazartesi")
	case time.Tuesday:
		fmt.Println("Salı")
	}
}

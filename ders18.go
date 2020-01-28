package main

import "fmt"

//structlar class'ların yerine geçer içindekiler field olarak adlandırılır
//type keyworld'ü ile kullanılır

type Human struct {
	FirstName string
	LastName  string
	Age       int
}

//structlarda constructor metotları yok
func newHuman() *Human {
	h := new(Human)
	return h
}

func main() {
	myHuman := Human{FirstName: "Aydemir Eray", LastName: "Han", Age: 22}
	// new ile nesne örneği oluşturma
	// myHuman2 := new(Human)
	// myHuman2.FirstName = "Aydemir Eray"

	//constructor ile
	// myHuman2 := newHuman()
	// myHuman2.FirstName = "Irmak"
	// myHuman2.LastName = "Han"
	// myHuman2.Age = 13
	// fmt.Println(myHuman2.FirstName)
	fmt.Println(myHuman.FirstName, myHuman.LastName, myHuman.Age)
}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type User struct {
	Id         int
	Email      string
	First_Name string
	Last_Name  string
	Avatar     string
}
type Page struct {
	Page        int
	Per_Page    int
	Total       int
	Total_Pages int
	Data        []User
}

func main() {
	data, err := http.Get("https://reqres.in/api/users?page=1")
	defer data.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	jsonData, err := ioutil.ReadAll(data.Body)
	if err != nil {
		log.Fatal(err)
	}
	var currentPage Page
	json.Unmarshal([]byte(jsonData), &currentPage)
	fmt.Printf("Page:%v \nPer Page:%v \nTotal:%v \nTotal Page:%v \n", currentPage.Page, currentPage.Per_Page, currentPage.Total, currentPage.Total_Pages)
	fmt.Println("\nKullanıcılar\n")
	for _,user := range currentPage.Data{
		fmt.Printf("Id:%v \nEmail:%v	\nFirst Name:%v	\nLast Name:%v \nAvatar:%v\n",user.Id,user.Email,user.First_Name,user.Last_Name,user.Avatar)
	}
}

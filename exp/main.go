package main

import (
	"html/template"
	"os"
)

type User struct {
	Name    string
	Age     int
	Work    bool
	DogName string
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	data := User{}
	// data := User{
	// 	Name: "John Smith",
	// 	DogName :=
	// }
	data.Name = "John Smith"
	data.Age = 35
	data.Work = true
	data.DogName = "Buddy"

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}

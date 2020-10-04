package main

import (
	"html/template"
	"os"
)

type User struct {
	Name  string
	Age   int
	Dog   Dog
	Slice []string
}

type Dog struct {
	Name string
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	data := User{
		Name: "John Smith",
		Age:  35,
		Dog: Dog{
			Name: "Jarsy",
		},
		Slice: []string{"qwe", "adf"},
	}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

}

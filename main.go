package main

import (
	"fmt"

	"github.com/hgko1207/learngo/mydict"
)


func main() {
	dictionary := mydict.Dictionary{"first": "First word"}
	word := "hello"
	definition := "Greeting"
	//definition, err := dictionary.Search("first")
	err := dictionary.Add(word, definition)
	if err != nil {
		fmt.Println(err)
	}

	hello, _ := dictionary.Search(word)
	fmt.Println("found:", word, "definition:", hello)

	err2 := dictionary.Add(word, definition)
	if err2 != nil {
		fmt.Println(err2)
	}
}
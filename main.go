package main

import (
	"fmt"

	"github.com/hgko1207/learngo/mydict"
)


func main() {
	dictionary := mydict.Dictionary{"first": "First word"}
	definition, err := dictionary.Search("first")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
}
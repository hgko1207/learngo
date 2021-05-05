package main

import (
	"fmt"
	"strings"
)


func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	name := "hgko"
	fmt.Println(name)

	fmt.Println(multiply(2, 2))

	totalLength, upperName := lenAndUpper("hgko")
	fmt.Println(totalLength, upperName)

	repeatMe("hgko", "ko", "eden")
}
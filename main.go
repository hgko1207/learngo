package main

import "fmt"

func canIDrink(age int) bool {
	if koreaAge := age + 2; koreaAge < 20 {
		return false
	}
	return true
}

func main() {
	fmt.Println((canIDrink(20)))
}
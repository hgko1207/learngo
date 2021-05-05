// Functions part One

package example

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

// naked return
func lenAndUpper2(name string) (length int, uppercase string) {
	// 'derfer' function이 끝난 후에 실행되는 코드
	defer fmt.Println("I'm done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

// func main() {
// 	fmt.Println(multiply(2, 2))

// 	totalLength, upperName := lenAndUpper("hgko")
// 	fmt.Println(totalLength, upperName)

// 	repeatMe("hgko", "ko", "eden")
// }
# Go

[노마드 코더](https://nomadcoders.co/) 에서 무료인 **쉽고 빠른 Go 시작하기** 강의 내용을 코딩하며 정리하였습니다.

## 시작

1. mkdir github.com/hgko1207
2. cd github.com/hgko1207
3. mkdir learngo
4. cd learngo
5. go mod init github.com/hgko1207/learngo

## 코딩

- main.go 는 컴파일을 위한 파일
- function -> func
- export 함수를 만들 경우 함수명 맨 앞에 대문자로 작성

## 상수와 변수

```go
// Constants
const name string = "hgko"
```

```go
// Variables
func main() {
    var name string = "hgko"

    // 변수를 아래와 같이 축약형으로 사용 가능
    // 축약형은 함수 안에서만 동작
    name := "hgko"
}
```

## 함수

```go
// naked return
func lenAndUpper(name string) (length int, uppercase string) {
	// 'derfer' function이 끝난 후에 실행되는 코드
	defer fmt.Println("I'm done")

	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}
```

# 참고

- https://golang.org/

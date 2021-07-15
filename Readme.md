# Go

[노마드 코더](https://nomadcoders.co/) 에서 무료인 **쉽고 빠른 Go 시작하기** 강의 내용을 코딩하며 정리하였습니다.

## 시작

https://go.dev/ 페이지에서 Go 를 다운로드 합니다. 다운로드 완료 시 설치파일을 실행하여 설치합니다. 그리고 폴더를 아래와 같이 생성합니다.

1. mkdir github.com/hgko1207
2. cd github.com/hgko1207
3. mkdir learngo
4. cd learngo
5. go mod init github.com/hgko1207/learngo

## 코딩

- main.go 는 컴파일을 위한 파일
- function -> func
- export 함수를 만들 경우 함수명 맨 앞에 대문자로 작성
- nil -> null을 뜻함

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

// Go가 내부적으로 호출하는 method를 사용하는 방법
func (a Account) String() string {
    return fmt.Sprint(a.Owner(), "'s account.\nHas: ", a.Balance())
}
```

## Map 사용

```go
// empty map을 초기화하고 싶을때
// map 끝에 중괄호({})를 작성하여 초기화 해야 함
var results = map[string]string{}

// 또는 make 사용
var results = make(map[string]string)

results["hello"] = "Hello"
```

## Goroutines

Goroutines이란 기본적으로 다른 함수와 동시에 실행시키는 함수

## 에러

errors.New() 를 사용하여 에러를 정의하고 호출한다.

```go
var err = errors.New("error content")
```

## 라이브러리

- go lang std library
- https://golang.org/pkg/

# 참고

- https://golang.org/

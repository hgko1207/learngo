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

## For 문

```go
people := [2]string{"hgko", "haeun"}
for _, person := range people {
    fmt.Println(person);
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

## 배열 선언

```go
func main() {
    data := [2]string{"test1", "test2"}
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

- Goroutines이란 기본적으로 다른 함수와 동시에 실행시키는 함수이다.
- Goroutines는 프로그램이 작동하는 동안(메인함수가 실행하는 동안)만 유효하다.

```go
// 함수 앞에 go를 붙인다.
main() {
    go count()
    count()
}

func count() {
    for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

// 이 경우는 main함수가 바로 종료되기 때문에 동작하지 않는다.
main() {
    go count()
    go count()
}
```

## Channels

- Channel은 Goroutines이랑 메인함수 사이에 정보를 전달하기 위한 방법이다.
- 채널로부터 뭔가를 받을 때 메인 함수가 어떤 결과가 올때까지 기다린다.
- 채널의 타입은 `chan` 이다.
- 채널은 데이터를 받을 순 없고 보낼 수만 있도록 작성할 수 있다.(send-only)

```go
func main() {
    c := make(chan bool)
    go channelTest(c)
    result := <- c
	fmt.Println(result)
}

func channelTest(c chan bool) {
    time.Sleep(time.Second * 5)
    c <- true
}

// send-only
func sendOnlyTest(c chan<- bool) {
    c <- true
}
```

## Error

- errors.New() 를 사용하여 에러를 정의하고 호출한다.

```go
var err = errors.New("error content")

func main() {
    fmt.PrintLn(errTest())
}

func errTest() error {
    return err
}
```

## 사용 라이브러리

```go
strings.Join()
strings.Fields()
strings.TrimSpace()
```

## goquery

- HTML을 navigate하고 필요한 걸 HTML에서 찾아내기 위해서 사용
- JQuery 와 유사한 go를 위한 라이브러리
- 참고 : https://github.com/PuerkitoBio/goquery

```bash
# 설치
$ go get github.com/PuerkitoBio/goquery
```

## Echo

설치

```bash
go get github.com/labstack/echo
```

echo 에 url 을 설정 할 수 있다.
아래 코드를 실행하고 http://localhost:1323 로 접속을 하게 되면 `Hello World!` 텍스트가 보인다.

```go
package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func handleHome(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.Logger.Fatal(e.Start(":1323"))
}
```

문자열 대신에 html 파일을 생성하여 html template을 전달 할 수 있다.

```html
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Go Home</title>
  </head>
  <body>
    <h1>Go Home</h1>
  </body>
</html>
```

home.html 파일을 생성하고 위의 코드를 복사한다. 그리고 `handleHome` 함수를 수정한다.

```go
func handleHome(c echo.Context) error {
	return c.File("home.html")
}
```

서버를 실행하면 Go Home 텍스트가 보인다.

# 참고

- [go lang std library](https://golang.org/pkg/)
- https://golang.org/
- [Online CSV Editor and Viewer](http://convertcsv.com/csv-viewer-editor.htm)

# Go

## 시작

1. cd github.com/hgko1207
2. mkdir learngo
3. cd learngo
4. go mod init github.com/hgko1207/learngo

## 코딩

- main.go 는 컴파일을 위한 파일이다.
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

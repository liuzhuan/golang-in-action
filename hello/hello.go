package hello

import "fmt"

func main() {
	fmt.Println("Hello World!");
}

// if `go run hello.go`, it will complain: "go run: cannot run non-main package"
// if `go build hello.go`, nothing will happen

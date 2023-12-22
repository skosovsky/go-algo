package main

import (
	"fmt"
)

func main() {
	fmt.Println(greet("Bob"))
}

func greet(name string) string {
	hello := fmt.Sprintf("Hello, %s how are you doing today?", name)
	return hello
}

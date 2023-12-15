package main

import (
	"fmt"
	"log"
)

func main() {
	var a, b int
	_, err := fmt.Scan(&a, &b) // standard input
	if err != nil {
		log.Println(err)
	}
	fmt.Println(a + b) // standard output
}

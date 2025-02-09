package main

import (
	"fmt"
	"os"

	"you_are_l/handlers"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <domain>")
		os.Exit(1)
	}

	domain := os.Args[1]

	result, err := handler.Whois(domain)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(result)
}

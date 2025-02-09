package main

import (
	"fmt"
	"os"

	"you_are_l/handlers"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <full|brief> <domain>")
		os.Exit(1)
	}

	command := os.Args[1]
	domain := os.Args[2]

	result, err := handler.Whois(domain)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	if command == "brief" {
		result = handler.ExtractBrief(result)
	}

	fmt.Println(result)
}

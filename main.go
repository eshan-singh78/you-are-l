package main

import (
	"fmt"
	"os"

	"you_are_l/handlers"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage:")
		fmt.Println("  WHOIS:         go run main.go whois <full|brief> <domain>")
		fmt.Println("  Unredirector:  go run main.go unredirect <short-url>")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "whois":
		if len(os.Args) < 4 {
			fmt.Println("Usage: go run main.go whois <full|brief> <domain>")
			os.Exit(1)
		}
		mode := os.Args[2]
		domain := os.Args[3]

		result, err := handler.Whois(domain)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		if mode == "brief" {
			result = handler.ExtractBrief(result)
		}

		fmt.Println(result)

	case "unredirect":
		shortURL := os.Args[2]
		finalURL, err := handler.Unredirect(shortURL)

		if err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			fmt.Printf("Final URL: %s\n", finalURL)
		}

	default:
		fmt.Println("Invalid command. Use 'whois' or 'unredirect'.")
		os.Exit(1)
	}
}

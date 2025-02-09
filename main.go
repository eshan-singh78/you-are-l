package main

import (
	"fmt"
	"os"

	"you_are_l/handlers"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage:")
		fmt.Println("  WHOIS:          go run main.go whois <full|brief> <domain>")
		fmt.Println("  Unredirector:   go run main.go unredirect <short-url>")
		fmt.Println("  DNS Lookup:     go run main.go dns <domain>")
		fmt.Println("  Content Fetch:  go run main.go fetch <url>")
		fmt.Println("  SSL Info:       go run main.go ssl <domain>")
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

	case "dns":
		domain := os.Args[2]
		records, err := handler.GetDNSRecords(domain)

		if err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			fmt.Println(records)
		}

	case "fetch":
		url := os.Args[2]
		info, err := handler.FetchContentInfo(url)

		if err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			fmt.Println(info)
		}

	case "ssl":
		domain := os.Args[2] // ✅ FIXED
		result, err := handler.FetchSSLCertificate(domain)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		fmt.Println(result)

	default:
		fmt.Println("Invalid command. Use 'whois', 'unredirect', 'dns', 'fetch', or 'ssl'.")
		os.Exit(1)
	}
}

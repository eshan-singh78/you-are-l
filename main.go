package main

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"you_are_l/handlers"
)

var rootCmd = &cobra.Command{
	Use:   "you-are-l",
	Short: "A CLI tool for domain & URL intelligence",
	Long:  "you-are-l is a command-line tool for performing WHOIS lookups, DNS queries, SSL certificate checks, URL unredirection, and content fetching.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use 'you-are-l --help' for available commands.")
	},
}

var whoisCmd = &cobra.Command{
	Use:   "whois [full|brief] <domain>",
	Short: "Perform a WHOIS lookup",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		mode, domain := args[0], args[1]
		result, err := handler.Whois(domain)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		if mode == "brief" {
			result = handler.ExtractBrief(result)
		}
		fmt.Println(result)
	},
}

var unredirectCmd = &cobra.Command{
	Use:   "unredirect <short-url>",
	Short: "Find the final destination of a shortened URL",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		finalURL, err := handler.Unredirect(args[0])
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		fmt.Println("Final URL:", finalURL)
	},
}

var dnsCmd = &cobra.Command{
	Use:   "dns <domain>",
	Short: "Retrieve DNS records for a domain",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		records, err := handler.GetDNSRecords(args[0])
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		fmt.Println(records)
	},
}

var fetchCmd = &cobra.Command{
	Use:   "fetch <url>",
	Short: "Fetch metadata and headers from a URL",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		info, err := handler.FetchContentInfo(args[0])
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		fmt.Println(info)
	},
}

var sslCmd = &cobra.Command{
	Use:   "ssl <domain>",
	Short: "Retrieve SSL certificate details",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		result, err := handler.FetchSSLCertificate(args[0])
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		fmt.Println(result)
	},
}

func main() {
	rootCmd.AddCommand(whoisCmd, unredirectCmd, dnsCmd, fetchCmd, sslCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

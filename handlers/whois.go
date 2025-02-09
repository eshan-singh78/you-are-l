package handler

import (
	"fmt"
	"io"
	"net"
	"strings"
	"time"
)

var whoisServers = map[string]string{
	"com":  "whois.verisign-grs.com",
	"net":  "whois.verisign-grs.com",
	"org":  "whois.pir.org",
	"gov":  "whois.dotgov.gov",
	"edu":  "whois.educause.edu",
	"info": "whois.afilias.net",
	"io":   "whois.nic.io",
}

func findWHOISServer(domain string) string {
	parts := strings.Split(domain, ".")
	if len(parts) < 2 {
		return "whois.iana.org"
	}

	tld := parts[len(parts)-1]
	if server, exists := whoisServers[tld]; exists {
		return server
	}
	return "whois.iana.org"
}

func Whois(domain string) (string, error) {
	server := findWHOISServer(domain)

	conn, err := net.DialTimeout("tcp", server+":43", 10*time.Second)
	if err != nil {
		return "", fmt.Errorf("failed to connect to WHOIS server: %v", err)
	}
	defer conn.Close()

	fmt.Fprintf(conn, domain+"\r\n")

	response, err := io.ReadAll(conn)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %v", err)
	}

	return string(response), nil
}

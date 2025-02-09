package handler

import (
	"fmt"
	"net"
	"strings"
)

func GetDNSRecords(domain string) (string, error) {
	var result []string

	aRecords, _ := net.LookupHost(domain)
	if len(aRecords) > 0 {
		result = append(result, "A Records:")
		result = append(result, aRecords...)
	}

	aaaaRecords, _ := net.LookupIP(domain)
	if len(aaaaRecords) > 0 {
		result = append(result, "AAAA Records:")
		for _, ip := range aaaaRecords {
			result = append(result, ip.String())
		}
	}

	soaRecord, err := net.LookupNS(domain)
	if err == nil && len(soaRecord) > 0 {
		result = append(result, "SOA Record:")
		for _, ns := range soaRecord {
			result = append(result, ns.Host)
		}
	}

	if len(result) == 0 {
		return "", fmt.Errorf("no DNS records found for %s", domain)
	}

	return strings.Join(result, "\n"), nil
}

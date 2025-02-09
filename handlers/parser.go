package handler

import (
	"regexp"
	"strings"
)

func ExtractBrief(whoisData string) string {
	lines := strings.Split(whoisData, "\n")
	var extracted []string

	patterns := []*regexp.Regexp{
		regexp.MustCompile(`(?i)^Domain Name:\s*`),
		regexp.MustCompile(`(?i)^Registry Domain ID:\s*`),
		regexp.MustCompile(`(?i)^Registrar WHOIS Server:\s*`),
		regexp.MustCompile(`(?i)^Registrar URL:\s*`),
		regexp.MustCompile(`(?i)^Updated Date:\s*`),
		regexp.MustCompile(`(?i)^Creation Date:\s*`),
		regexp.MustCompile(`(?i)^Registry Expiry Date:\s*`),
		regexp.MustCompile(`(?i)^Registrar:\s*`),
		regexp.MustCompile(`(?i)^Registrar IANA ID:\s*`),
		regexp.MustCompile(`(?i)^Registrar Abuse Contact Email:\s*`),
		regexp.MustCompile(`(?i)^Registrar Abuse Contact Phone:\s*`),
		regexp.MustCompile(`(?i)^Domain Status:\s*`),
		regexp.MustCompile(`(?i)^Name Server:\s*`),
		regexp.MustCompile(`(?i)^DNSSEC:\s*`),
	}

	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)

		for _, pattern := range patterns {
			if pattern.MatchString(trimmedLine) {
				extracted = append(extracted, trimmedLine)
				break
			}
		}
	}

	if len(extracted) == 0 {
		return "No matching WHOIS data found."
	}

	return strings.Join(extracted, "\n")
}

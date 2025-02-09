package handler

import (
	"crypto/md5"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"strings"
)

func FetchSSLCertificate(domain string) (string, error) {
	if !strings.HasPrefix(domain, "https://") && !strings.HasPrefix(domain, "http://") {
		domain = "https://" + domain
	}

	domain = strings.TrimPrefix(domain, "https://")
	domain = strings.TrimPrefix(domain, "http://")

	conn, err := tls.Dial("tcp", domain+":443", &tls.Config{})
	if err != nil {
		return "", fmt.Errorf("failed to connect to SSL server: %v", err)
	}
	defer conn.Close()

	cert := conn.ConnectionState().PeerCertificates[0]

	serialNumber := fmt.Sprintf("%X", cert.SerialNumber)
	fingerprint := md5.Sum(cert.Raw)
	fingerprintStr := formatFingerprint(hex.EncodeToString(fingerprint[:]))

	var extendedKeyUsages []string
	for _, eku := range cert.ExtKeyUsage {
		switch eku {
		case x509.ExtKeyUsageServerAuth:
			extendedKeyUsages = append(extendedKeyUsages, "TLS Web Server Authentication")
		case x509.ExtKeyUsageClientAuth:
			extendedKeyUsages = append(extendedKeyUsages, "TLS Web Client Authentication")
		}
	}

	result := []string{
		fmt.Sprintf("Subject: %s", cert.Subject.CommonName),
		fmt.Sprintf("Issuer: %s", cert.Issuer.CommonName),
		fmt.Sprintf("Expires: %s", cert.NotAfter.Format("2 January 2006")),
		fmt.Sprintf("Renewed: %s", cert.NotBefore.Format("2 January 2006")),
		fmt.Sprintf("Serial Num: %s", serialNumber),
		fmt.Sprintf("Fingerprint: %s", fingerprintStr),
		fmt.Sprintf("Extended Key Usage: %s", strings.Join(extendedKeyUsages, ", ")),
	}

	return strings.Join(result, "\n"), nil
}

func formatFingerprint(fingerprint string) string {
	var formatted []string
	for i := 0; i < len(fingerprint); i += 2 {
		formatted = append(formatted, fingerprint[i:i+2])
	}
	return strings.ToUpper(strings.Join(formatted, ":"))
}

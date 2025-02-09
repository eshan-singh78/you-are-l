package handler

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"golang.org/x/net/html"
)

func FetchContentInfo(url string) (string, error) {
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "https://" + url
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to fetch URL: %v", err)
	}
	defer resp.Body.Close()

	var metaTitle, metaDesc string

	tokenizer := html.NewTokenizer(resp.Body)
	for {
		tokenType := tokenizer.Next()
		switch tokenType {
		case html.ErrorToken:
			return formatResponse(url, resp, metaTitle, metaDesc), nil
		case html.StartTagToken:
			token := tokenizer.Token()
			if token.Data == "title" {
				tokenizer.Next()
				metaTitle = tokenizer.Token().Data
			} else if token.Data == "meta" {
				nameAttr, contentAttr := "", ""
				for _, attr := range token.Attr {
					if attr.Key == "name" || attr.Key == "property" {
						nameAttr = strings.ToLower(attr.Val)
					}
					if attr.Key == "content" {
						contentAttr = attr.Val
					}
				}
				if nameAttr == "description" {
					metaDesc = contentAttr
				}
			}
		}
	}
}

func formatResponse(url string, resp *http.Response, title, desc string) string {
	var result []string

	result = append(result, fmt.Sprintf("Fetched URL: %s", url))
	result = append(result, fmt.Sprintf("Status Code: %d", resp.StatusCode))

	if title != "" {
		result = append(result, fmt.Sprintf("Title: %s", title))
	}
	if desc != "" {
		result = append(result, fmt.Sprintf("Description: %s", desc))
	}

	result = append(result, "\nHeaders:")
	for key, values := range resp.Header {
		result = append(result, fmt.Sprintf("%s: %s", key, strings.Join(values, ", ")))
	}

	return strings.Join(result, "\n")
}

package handler

import (
	"fmt"
	"net/http"
	"time"
)

func Unredirect(shortURL string) (string, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if len(via) >= 10 {
				return fmt.Errorf("too many redirects")
			}
			return nil
		},
	}

	resp, err := client.Get(shortURL)
	if err != nil {
		return "", fmt.Errorf("failed to resolve URL: %v", err)
	}
	defer resp.Body.Close()

	return resp.Request.URL.String(), nil
}

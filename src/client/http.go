package client

import (
	"net/http"
	"time"
)

func NewClient(checkRedirect func(*http.Request, []*http.Request) error, timeout time.Duration) *http.Client {
	return &http.Client{
		CheckRedirect: checkRedirect,
		Timeout:       timeout,
	}
}

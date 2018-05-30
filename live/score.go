package live

import (
	"client"
	"fmt"
	"io"
	"net/http"
	"sports"
	"strings"
	"time"
)

var footballUrl = "https://www.mbet.com/zh/live/" + sports.SFootballID

func checkRedirect(req *http.Request, via []*http.Request) error {
	if len(via) > 10 {
		return fmt.Errorf("too many redirect\n")
	}
	if req.Response != nil {
		if req.Response.StatusCode == http.StatusFound {
			return fmt.Errorf("redirect 302\n")
		}
	}
	return nil
}

func HasFootball() (bool, io.Reader, error) {
	clt := client.NewClient(checkRedirect, 20*time.Second)
	req, err := http.NewRequest("GET", footballUrl, nil)
	if err != nil {
		return false, nil, fmt.Errorf("NewRequest: %v\n", err)
	}
	resp, err := clt.Do(req)
	if err != nil {
		if strings.Contains(err.Error(), "302") {
			return false, nil, nil
		} else {
			return false, nil, fmt.Errorf("http client Do: %v\n", err)
		}
	}

	if resp.StatusCode != http.StatusOK {
		return false, nil, nil
	}
	return true, resp.Body, nil
}

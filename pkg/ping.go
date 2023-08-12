package pkg

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync/atomic"
	"time"
)

func PingSite(site string) int {
	statusChan := make(chan int)

	u, err := url.Parse(site)
	if err != nil {
		fmt.Printf("could not parse url: %s\n", err)
		return http.StatusInternalServerError
	}

	var workers int32 = 4
	for i := int32(0); i < workers; i++ {
		go func() {
			defer func() {
				if atomic.AddInt32(&workers, -1) <= 0 {
					close(statusChan)
				}
			}()
			statusChan <- GetStatus(u)
		}()
	}

	finalStatus := http.StatusOK
	for s := range statusChan {
		fmt.Println(s)
		if finalStatus == http.StatusOK && s != http.StatusOK {
			finalStatus = s
		}
	}

	return finalStatus
}

func GetStatus(u *url.URL) int {
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("%s://%s", "https", u.Host),
		bytes.NewReader([]byte{}),
	)
	if err != nil {
		fmt.Printf("could not create request: %s\n", err)
		return http.StatusInternalServerError
	}

	req.Header.Set("Content-Type", "application/json")

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Do(req)
	if res != nil {
		defer res.Body.Close()
	}
	if err != nil {
		fmt.Printf("request failed: %s\n", err)
		return http.StatusInternalServerError
	}

	_, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("could not read response: %s\n", err)
		return http.StatusInternalServerError
	}

	// fmt.Printf("response: %s\n", string(b))

	return http.StatusOK
}

package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"
	"sync"
)

type status int

const (
	Success status = iota
	PermanentRedirect
	TemporaryRedirect
	Error
)

type result struct {
	Code     int
	Location string
	Status   status
	URL      string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	concurrency := flag.Int("num", 5, "Concurrency rate")
	wg := &sync.WaitGroup{}
	queue := make(chan string)
	results := make(chan result)

	flag.Parse()
	wg.Add(*concurrency)

	go func() {
		for result := range results {
			if result.Status == TemporaryRedirect || result.Status == PermanentRedirect {
				fmt.Printf("StatusCode:%d URL:%s Location:%s\n", result.Code, result.URL, result.Location)
			} else {
				fmt.Printf("StatusCode:%d URL:%s\n", result.Code, result.URL)
			}
		}
	}()

	for i := 0; i < *concurrency; i++ {
		go func() {
			defer wg.Done()

			for url := range queue {
				request(url, results)
			}
		}()
	}

	for scanner.Scan() {
		queue <- scanner.Text()
	}

	close(queue)

	wg.Wait()
}

func request(uri string, results chan result) {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	res, err := client.Head(uri)

	switch {
	case err != nil:
		results <- result{
			Status: Error,
			URL:    uri,
		}
	case res.StatusCode >= 200 && res.StatusCode < 300:
		results <- result{
			Status: Success,
			Code:   res.StatusCode,
			URL:    uri,
		}
	case res.StatusCode == 301:
		results <- result{
			Status:   PermanentRedirect,
			Code:     res.StatusCode,
			Location: res.Header.Get("location"),
			URL:      uri,
		}
	case res.StatusCode == 302 || res.StatusCode == 303 || res.StatusCode == 307:
		results <- result{
			Status:   TemporaryRedirect,
			Code:     res.StatusCode,
			Location: res.Header.Get("location"),
			URL:      uri,
		}
	default:
		results <- result{
			Status: Error,
			Code:   res.StatusCode,
			URL:    uri,
		}
	}
}

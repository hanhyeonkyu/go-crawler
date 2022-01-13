package urlcheck

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

var errReqFailed = errors.New("request failed")
var urls = []string{
	"https://www.airbnb.com/",
	"https://www.google.com/",
	"https://www.amazon.com/",
	"https://www.raddit.com/",
	"https://www.google.com/",
	"https://soundcloud.com/",
	"https://www.facebook.com/",
	"https://www.instagram.com/",
	"https://academy.nomadcoders.co/",
}

func UrlChecker() {
	var results = make(map[string]string)
	for _, url := range urls {
		result := "OK"
		err := hitURL(url)
		if err != nil {
			result = "FAILED"
		}
		results[url] = result
	}
	fmt.Println(results)
}

func hitURL(url string) error {
	fmt.Println("CHECKING URL:", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		return errReqFailed
	}
	return nil
}

type result struct {
	url    string
	status string
}

func UrlCheckerWithRoutine() {
	c := make(chan result)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.raddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}
	for _, url := range urls {
		go hitURLWithRoutine(url, c)
	}
	for i := 0; i < len(urls); i++ {
		fmt.Println("CHECKING URL:", <-c)
	}
}

func hitURLWithRoutine(url string, c chan<- result) {
	time.Sleep(time.Second * 1)
	resp, err := http.Get(url)
	status := "SUCCESS"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
		c <- result{url, status}
	}
	c <- result{url, status}
}

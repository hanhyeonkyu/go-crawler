package urlcheck

import (
	"errors"
	"fmt"
	"net/http"
)

var errReqFailed = errors.New("request failed")

func UrlChecker() {
	var results = make(map[string]string)
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

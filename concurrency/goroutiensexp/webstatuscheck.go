package goroutiensexp

import (
	"fmt"
	"net/http"
)

func WebStatus() {
	fmt.Println("go rouotines")
	weblinks := []string{
		"http://google.com",
		"http://facebook.com",
		"http://twitter.com",
		"http://linkedin.com",
	}

	for idx, link := range weblinks {
		fmt.Println("web link index: ", idx, "----", "web link: ", link)
		CheckStatus(link)

	}
}

func CheckStatus(url string) {
	_, err := http.Get((url)) // Blocking program execution until response came back.

	if err != nil {
		fmt.Println("website might down: " + url)
		return
	}

	fmt.Println("Website is fine: " + url)
}

// This program run in serial, check the each website one by one.
// this adds some delay in program to run, can impore by adding parallesim using goroutines.

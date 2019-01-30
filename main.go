package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func doEvery(d time.Duration, f func(string, time.Time)) {
	urls := os.Getenv("URLS")
	if urls == "" {
		fmt.Println("URLS environment variable must be defined")
		os.Exit(1)
	}
	fmt.Println("Starting check on the following urls:", urls)
	for x := range time.Tick(d) {
		for _, url := range strings.Split(urls, ",") {
			f(url, x)
		}
	}
}

func checkURL(url string, t time.Time) {
	_, err := net.DialTimeout("tcp", url, 6*time.Second)
	if err != nil {
		log.Println("Site unreachable, error: ", err)
		return
	}
	fmt.Printf("%v: %s works!\n", t, url)
}

func main() {
	doEvery(2000*time.Millisecond, checkURL)
}

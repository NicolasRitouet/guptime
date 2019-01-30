package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"runtime"
	"strings"
	"time"
)

func doEvery(d time.Duration, f func(string, time.Time)) {
	urls := os.Getenv("URLS")
	if urls == "" {
		fmt.Println("URLS environment variable must be defined")
		os.Exit(1)
	}
	fmt.Println(urls)
	for x := range time.Tick(d) {
		for _, url := range strings.Split(urls, ",") {
			f(url, x)
		}
	}
}

func check(url string, t time.Time) {
	// printMemUsage()
	_, err := net.DialTimeout("tcp", url, 6*time.Second)
	if err != nil {
		log.Println("Site unreachable, error: ", err)
		return
	}
	fmt.Printf("%v: %s works!\n", t, url)
}

func printMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func main() {
	doEvery(2000*time.Millisecond, check)
}

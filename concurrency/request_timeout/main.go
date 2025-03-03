package main

import (
	"fmt"
	"time"

	"math/rand"
)

func doRequest() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	time.Sleep(time.Duration(r.Intn(3)) * time.Second) // simulate lasting request
}

func handleRequest() {
	defer func(start time.Time) {
		fmt.Printf("request took %f seconds\n", time.Since(start).Seconds())
	}(time.Now())

	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop() // required before Go 1.23 to stop writing into channel time.Time.C

	timer := time.NewTimer(time.Second)
	defer timer.Stop() // required before Go 1.23 to stop writing into channel time.Time.C

	done := make(chan struct{})

	go func() {
		doRequest()
		done <- struct{}{}
	}()

	for {
		select {
		case <-done:
			fmt.Println("request done")
			return
		case <-timer.C:
			fmt.Println("request timeout")
			return
		case <-ticker.C:
			fmt.Println("request handling")
		}
	}
}

func main() {
	handleRequest()
}

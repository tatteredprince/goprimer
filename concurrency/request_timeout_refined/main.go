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

	ticker := time.Tick(500 * time.Millisecond)

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
		case <-time.After(time.Second):
			fmt.Println("request timeout")
			return
		case <-ticker:
			fmt.Println("request handling")
		}
	}
}

func main() {
	handleRequest()
}

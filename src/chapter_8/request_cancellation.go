package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func request(ctx context.Context, url string) string {
	rand.Seed(time.Now().UnixNano())

	select {
	// pretend to make a http request
	case <-time.After(time.Duration(rand.Intn(3))):
		return url
	case <-ctx.Done():
		fmt.Printf("\n\naaaaaaa %+v\n\n", "cancelling: "+url)
		return ""
	}
}

func getFastestResponse() string {
	responses := make(chan string, 3)

	urls := []string{"asia.gopl.io", "americas.gopl.io", "europe.gopl.io"}

	ctx, cancel := context.WithCancel(context.Background())

	for _, url := range urls {
		url := url
		go func() {
			responses <- request(ctx, url)
		}()
	}

	response := <-responses

	cancel()

	time.Sleep(2 * time.Second)

	return response
}

func main() {
	fmt.Println(getFastestResponse())
}

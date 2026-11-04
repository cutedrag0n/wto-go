package main

import (
	"context"
	"fmt"
	wto "github.com/cutedrag0n/wto-go"
	"github.com/cutedrag0n/wto-go/qr"
	"log"
	"os"
)

func main() {
	limiter := wto.NewRateLimiter()
	first, err := wto.NewClient(os.Getenv("WTO_API_KEY"), wto.WithRateLimiter(limiter))
	if err != nil {
		log.Fatal(err)
	}
	second, err := wto.NewClient(os.Getenv("WTO_API_KEY"), wto.WithRateLimiter(limiter))
	if err != nil {
		log.Fatal(err)
	}
	for _, client := range []*wto.Client{first, second} {
		v, err := client.QR.Members(context.Background(), qr.MembersParams{})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Members: %d\n", len(v.Data))
	}
}

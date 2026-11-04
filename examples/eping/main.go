package main

import (
	"context"
	"fmt"
	wto "github.com/cutedrag0n/wto-go"
	"github.com/cutedrag0n/wto-go/eping"
	"log"
	"os"
)

func main() {
	c, err := wto.NewClient(os.Getenv("WTO_API_KEY"))
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	v, err := c.EPing.SearchNotifications(ctx, eping.SearchNotificationsParams{Language: "1"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Notifications: %d\n", len(v.Items))
}

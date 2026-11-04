package main

import (
	"context"
	"fmt"
	"log"
	"os"

	wto "github.com/cutedrag0n/wto-go"
	"github.com/cutedrag0n/wto-go/timeseries"
)

func main() {
	c, err := wto.NewClient(os.Getenv("WTO_API_KEY"))
	if err != nil {
		log.Fatal(err)
	}
	max := 10
	data, err := c.Timeseries.Data(context.Background(), timeseries.DataParams{Indicator: "ITS_MTV_AX", Max: &max})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Retrieved %d datapoints\n", len(data))
}

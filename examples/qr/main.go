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
	c, err := wto.NewClient(os.Getenv("WTO_API_KEY"))
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	v, err := c.QR.Products(ctx, qr.ProductsParams{HSVersion: "h6", Code: "0101"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Products: %d\n", len(v.Data))
}

package main

import (
	"context"
	"fmt"
	wto "github.com/cutedrag0n/wto-go"
	"github.com/cutedrag0n/wto-go/tfad"
	"log"
	"os"
)

func main() {
	c, err := wto.NewClient(os.Getenv("WTO_API_KEY"))
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	v, err := c.TFAD.Transparency(ctx, tfad.TransparencyParams{Countries: []string{"840"}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(v))
}

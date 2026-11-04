package wto_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"

	wto "github.com/cutedrag0n/wto-go"
	"github.com/cutedrag0n/wto-go/timeseries"
)

func ExampleNewClient() {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `[{"code":"A","name":"Annual"}]`)
	}))
	defer server.Close()
	client, err := wto.NewClient("example-key", wto.WithBaseURL(server.URL))
	if err != nil {
		fmt.Println(err)
		return
	}
	frequencies, err := client.Timeseries.Frequencies(context.Background(), timeseries.FrequenciesParams{})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(*frequencies[0].Code)
	// Output: A
}

package tfad_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	wto "github.com/cutedrag0n/wto-go"
	"github.com/cutedrag0n/wto-go/tfad"
)

func ptr(v int) *int { return &v }
func TestOfficialOperations(t *testing.T) {
	t.Run("Transparency", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/tfad/transparency/procedures_contacts_single_window" || r.Method != "GET" {
				t.Errorf("unexpected request %s %s", r.Method, r.URL.Path)
			}
			if r.Header.Get("Ocp-Apim-Subscription-Key") != "test-key" {
				t.Error("missing key")
			}
			w.Write([]byte(`{"members":[]}`))
		}))
		defer server.Close()
		c, err := wto.NewClient("test-key", wto.WithBaseURL(server.URL))
		if err != nil {
			t.Fatal(err)
		}
		if _, err := c.TFAD.Transparency(context.Background(), tfad.TransparencyParams{}); err != nil {
			t.Fatal(err)
		}
	})
}

func TestCountryFilters(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		v := r.URL.Query()["countries[]"]
		if len(v) != 2 || v[0] != "004" || v[1] != "008" {
			t.Error(v)
		}
		w.Write([]byte(`{"members":[]}`))
	}))
	defer server.Close()
	c, err := wto.NewClient("key", wto.WithBaseURL(server.URL))
	if err != nil {
		t.Fatal(err)
	}
	if _, err := c.TFAD.Transparency(context.Background(), tfad.TransparencyParams{Countries: []string{"004", "008"}}); err != nil {
		t.Fatal(err)
	}
}

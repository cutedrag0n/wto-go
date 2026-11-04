package eping_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	wto "github.com/cutedrag0n/wto-go"
	"github.com/cutedrag0n/wto-go/eping"
)

func ptr(v int) *int { return &v }
func TestOfficialOperations(t *testing.T) {
	t.Run("SearchNotifications", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/eping/notifications/search" || r.Method != "GET" {
				t.Errorf("unexpected request %s %s", r.Method, r.URL.Path)
			}
			if r.Header.Get("Ocp-Apim-Subscription-Key") != "test-key" {
				t.Error("missing key")
			}
			w.Write([]byte(`{"items":[],"totalCount":0}`))
		}))
		defer server.Close()
		c, err := wto.NewClient("test-key", wto.WithBaseURL(server.URL))
		if err != nil {
			t.Fatal(err)
		}
		if _, err := c.EPing.SearchNotifications(context.Background(), eping.SearchNotificationsParams{Language: "1"}); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("Members", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/eping/members" || r.Method != "GET" {
				t.Errorf("unexpected request %s %s", r.Method, r.URL.Path)
			}
			if r.Header.Get("Ocp-Apim-Subscription-Key") != "test-key" {
				t.Error("missing key")
			}
			w.Write([]byte(`[]`))
		}))
		defer server.Close()
		c, err := wto.NewClient("test-key", wto.WithBaseURL(server.URL))
		if err != nil {
			t.Fatal(err)
		}
		if _, err := c.EPing.Members(context.Background(), eping.MembersParams{Language: "1"}); err != nil {
			t.Fatal(err)
		}
	})
}

func TestSearchFilters(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		for k, want := range map[string]string{"language": "2", "domainIds": "1", "documentSymbol": "USA/1234", "distributionDateFrom": "2022-08-01", "distributionDateTo": "2022-09-01", "hs": "0102,0103", "ics": "01.030", "freeText": "food safety", "page": "2", "pageSize": "100"} {
			if q.Get(k) != want {
				t.Errorf("%s: %q", k, q.Get(k))
			}
		}
		if got := q["countryIds"]; len(got) != 2 || got[0] != "C004" || got[1] != "C840" {
			t.Error(got)
		}
		w.Write([]byte(`{"items":[],"totalCount":0}`))
	}))
	defer server.Close()
	c, err := wto.NewClient("key", wto.WithBaseURL(server.URL))
	if err != nil {
		t.Fatal(err)
	}
	_, err = c.EPing.SearchNotifications(context.Background(), eping.SearchNotificationsParams{
		Language: "2", DomainIDs: ptr(1), DocumentSymbol: "USA/1234", DistributionDateFrom: "2022-08-01", DistributionDateTo: "2022-09-01", CountryIDs: []string{"C004", "C840"}, HS: "0102,0103", ICS: "01.030", FreeText: "food safety", Page: ptr(2), PageSize: ptr(100),
	})
	if err != nil {
		t.Fatal(err)
	}
}

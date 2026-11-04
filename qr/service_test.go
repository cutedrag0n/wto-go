package qr_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	wto "github.com/cutedrag0n/wto-go"
	"github.com/cutedrag0n/wto-go/qr"
)

func ptr(v int) *int { return &v }
func TestOfficialOperations(t *testing.T) {
	t.Run("Members", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/qrs/members" || r.Method != "GET" {
				t.Errorf("unexpected request %s %s", r.Method, r.URL.Path)
			}
			if r.Header.Get("Ocp-Apim-Subscription-Key") != "test-key" {
				t.Error("missing key")
			}
			w.Write([]byte(`{"data":[],"meta":{"current_page":1,"total":0}}`))
		}))
		defer server.Close()
		c, err := wto.NewClient("test-key", wto.WithBaseURL(server.URL))
		if err != nil {
			t.Fatal(err)
		}
		if _, err := c.QR.Members(context.Background(), qr.MembersParams{}); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("Products", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/qrs/products" || r.Method != "GET" {
				t.Errorf("unexpected request %s %s", r.Method, r.URL.Path)
			}
			if r.Header.Get("Ocp-Apim-Subscription-Key") != "test-key" {
				t.Error("missing key")
			}
			w.Write([]byte(`{"data":[],"meta":{"current_page":1,"total":0}}`))
		}))
		defer server.Close()
		c, err := wto.NewClient("test-key", wto.WithBaseURL(server.URL))
		if err != nil {
			t.Fatal(err)
		}
		if _, err := c.QR.Products(context.Background(), qr.ProductsParams{HSVersion: "h6"}); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("Notifications", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/qrs/notifications" || r.Method != "GET" {
				t.Errorf("unexpected request %s %s", r.Method, r.URL.Path)
			}
			if r.Header.Get("Ocp-Apim-Subscription-Key") != "test-key" {
				t.Error("missing key")
			}
			w.Write([]byte(`{"data":[],"meta":{"current_page":1,"total":0}}`))
		}))
		defer server.Close()
		c, err := wto.NewClient("test-key", wto.WithBaseURL(server.URL))
		if err != nil {
			t.Fatal(err)
		}
		if _, err := c.QR.Notifications(context.Background(), qr.NotificationsParams{}); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("Restriction", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/qrs/qrs/14" || r.Method != "GET" {
				t.Errorf("unexpected request %s %s", r.Method, r.URL.Path)
			}
			if r.Header.Get("Ocp-Apim-Subscription-Key") != "test-key" {
				t.Error("missing key")
			}
			w.Write([]byte(`{"data":{}}`))
		}))
		defer server.Close()
		c, err := wto.NewClient("test-key", wto.WithBaseURL(server.URL))
		if err != nil {
			t.Fatal(err)
		}
		if _, err := c.QR.Restriction(context.Background(), qr.RestrictionParams{QRID: ptr(14)}); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("HSVersions", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/qrs/hs-versions" || r.Method != "GET" {
				t.Errorf("unexpected request %s %s", r.Method, r.URL.Path)
			}
			if r.Header.Get("Ocp-Apim-Subscription-Key") != "test-key" {
				t.Error("missing key")
			}
			w.Write([]byte(`{"data":[],"meta":{"current_page":1,"total":0}}`))
		}))
		defer server.Close()
		c, err := wto.NewClient("test-key", wto.WithBaseURL(server.URL))
		if err != nil {
			t.Fatal(err)
		}
		if _, err := c.QR.HSVersions(context.Background(), qr.HSVersionsParams{}); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("Restrictions", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/qrs/qrs" || r.Method != "GET" {
				t.Errorf("unexpected request %s %s", r.Method, r.URL.Path)
			}
			if r.Header.Get("Ocp-Apim-Subscription-Key") != "test-key" {
				t.Error("missing key")
			}
			w.Write([]byte(`{"data":[],"meta":{"current_page":1,"total":0}}`))
		}))
		defer server.Close()
		c, err := wto.NewClient("test-key", wto.WithBaseURL(server.URL))
		if err != nil {
			t.Fatal(err)
		}
		if _, err := c.QR.Restrictions(context.Background(), qr.RestrictionsParams{}); err != nil {
			t.Fatal(err)
		}
	})
}

func TestRestrictionFilters(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		for k, want := range map[string]string{"page": "2", "reporter_member_code": "C004", "in_force_only": "false", "year_of_entry_into_force": "2020", "product_codes": "0102,0103", "product_ids": "h6-0102,h7-0103"} {
			if q.Get(k) != want {
				t.Errorf("%s: %s", k, q.Get(k))
			}
		}
		w.Write([]byte(`{"data":[]}`))
	}))
	defer server.Close()
	c, err := wto.NewClient("key", wto.WithBaseURL(server.URL))
	if err != nil {
		t.Fatal(err)
	}
	no := false
	if _, err := c.QR.Restrictions(context.Background(), qr.RestrictionsParams{Page: ptr(2), ReporterMemberCode: "C004", InForceOnly: &no, YearOfEntryIntoForce: ptr(2020), ProductCodes: "0102,0103", ProductIDs: "h6-0102,h7-0103"}); err != nil {
		t.Fatal(err)
	}
}

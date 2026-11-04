package wto_test

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	wto "github.com/cutedrag0n/wto-go"
	"github.com/cutedrag0n/wto-go/qr"
)

func TestConfiguration(t *testing.T) {
	for _, opts := range [][]wto.Option{{wto.WithTimeout(0)}, {wto.WithRetry(-1)}, {wto.WithBaseURL("bad")}, {wto.WithBaseURL("https://user:pass@example.com")}, {wto.WithProxy("bad")}} {
		if _, err := wto.NewClient("key", opts...); err == nil {
			t.Fatal("invalid config accepted")
		}
	}
	if _, err := wto.NewClient(""); err == nil {
		t.Fatal("empty key")
	}
}
func TestSharedLimiter(t *testing.T) {
	var calls atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { calls.Add(1); w.Write([]byte(`{"data":[]}`)) }))
	defer server.Close()
	l := wto.NewRateLimiter()
	a, err := wto.NewClient("key", wto.WithRateLimiter(l), wto.WithBaseURL(server.URL))
	if err != nil {
		t.Fatal(err)
	}
	b, err := wto.NewClient("key", wto.WithRateLimiter(l), wto.WithBaseURL(server.URL))
	if err != nil {
		t.Fatal(err)
	}
	if _, err := a.QR.Members(context.Background(), qr.MembersParams{}); err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	defer cancel()
	if _, err := b.QR.Members(ctx, qr.MembersParams{}); !errors.Is(err, context.DeadlineExceeded) {
		t.Fatal(err)
	}
	if calls.Load() != 1 {
		t.Fatal(calls.Load())
	}
}

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

func TestCustomClientAndServiceURL(t *testing.T) {
	var calls atomic.Int32
	custom := &http.Client{Transport: roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		calls.Add(1)
		if r.URL.String() != "https://example.test/custom/members" || r.Header.Get("X-Test") != "value" {
			t.Error(r.URL, r.Header)
		}
		return &http.Response{StatusCode: 200, Header: http.Header{}, Body: io.NopCloser(strings.NewReader(`{"data":[]}`)), Request: r}, nil
	})}
	original := custom.Transport
	c, err := wto.NewClient("key", wto.WithHTTPClient(custom), wto.WithHeader("X-Test", "value"), wto.WithServiceBaseURL(wto.QR, "https://example.test/custom"))
	if err != nil {
		t.Fatal(err)
	}
	if _, err := c.QR.Members(context.Background(), qr.MembersParams{}); err != nil {
		t.Fatal(err)
	}
	if calls.Load() != 1 || custom.Transport == nil || reflect.ValueOf(custom.Transport).Pointer() != reflect.ValueOf(original).Pointer() || custom.CheckRedirect != nil {
		t.Fatal("custom client mutated")
	}
	if _, err := wto.NewClient("key", wto.WithHTTPClient(custom), wto.WithProxy("http://localhost:8080")); err == nil {
		t.Fatal("proxy accepted custom RoundTripper")
	}
}
func TestProxy(t *testing.T) {
	proxy := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Host != "example.test" {
			t.Error(r.URL)
		}
		w.Write([]byte(`{"data":[]}`))
	}))
	defer proxy.Close()
	c, err := wto.NewClient("key", wto.WithBaseURL("http://example.test"), wto.WithProxy(proxy.URL))
	if err != nil {
		t.Fatal(err)
	}
	if _, err := c.QR.Members(context.Background(), qr.MembersParams{}); err != nil {
		t.Fatal(err)
	}
}
func TestRedirectDoesNotLeakKey(t *testing.T) {
	var leaked atomic.Bool
	destination := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { leaked.Store(true) }))
	defer destination.Close()
	origin := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { http.Redirect(w, r, destination.URL, 302) }))
	defer origin.Close()
	c, err := wto.NewClient("secret", wto.WithBaseURL(origin.URL))
	if err != nil {
		t.Fatal(err)
	}
	_, err = c.QR.Members(context.Background(), qr.MembersParams{})
	var api *wto.APIError
	if !errors.As(err, &api) || api.StatusCode != 302 || leaked.Load() {
		t.Fatal("cross-origin redirect", err)
	}
}

package timeseries_test

import (
	"archive/zip"
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	wto "github.com/cutedrag0n/wto-go"
	"github.com/cutedrag0n/wto-go/timeseries"
)

func TestDataModesAndPost(t *testing.T) {
	for _, post := range []bool{false, true} {
		for _, mode := range []string{"full", "codes"} {
			t.Run(mode+map[bool]string{true: "POST", false: "GET"}[post], func(t *testing.T) {
				server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					var p timeseries.DataParams
					if post {
						if r.Method != "POST" {
							t.Error(r.Method)
						}
						if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
							t.Error(err)
						}
					} else {
						q := r.URL.Query()
						p.Indicator = q.Get("i")
						p.Headings = q.Get("head")
						p.Mode = q.Get("mode")
						if q.Get("off") != "0" || q.Get("spc") != "false" {
							t.Error(q)
						}
					}
					if p.Indicator != "ITS_MTV_AX" || p.Headings != "M" || p.Mode != mode {
						t.Error(p)
					}
					w.Write([]byte(`{"Dataset":[{"IndicatorCode":"ITS_MTV_AX","ReportingEconomyCode":"004","Value":0,"TextValue":null}]}`))
				}))
				defer server.Close()
				c, err := wto.NewClient("key", wto.WithBaseURL(server.URL))
				if err != nil {
					t.Fatal(err)
				}
				zero := 0
				no := false
				p := timeseries.DataParams{Indicator: "ITS_MTV_AX", Mode: mode, Offset: &zero, IncludeSubproducts: &no}
				fn := c.Timeseries.Data
				if post {
					fn = c.Timeseries.DataPost
				}
				rows, err := fn(context.Background(), p)
				if err != nil {
					t.Fatal(err)
				}
				if len(rows) != 1 || rows[0].ReportingEconomyCode == nil || *rows[0].ReportingEconomyCode != "004" || rows[0].Value == nil || *rows[0].Value != 0 || rows[0].TextValue != nil {
					t.Fatal(rows)
				}
			})
		}
	}
}
func TestDownload(t *testing.T) {
	var data bytes.Buffer
	z := zip.NewWriter(&data)
	f, err := z.Create("data.csv")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := f.Write([]byte("Indicator,Value\nx,1\n")); err != nil {
		t.Fatal(err)
	}
	if err := z.Close(); err != nil {
		t.Fatal(err)
	}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("fmt") != "csv" {
			t.Error(r.URL)
		}
		w.Write(data.Bytes())
	}))
	defer server.Close()
	c, err := wto.NewClient("key", wto.WithBaseURL(server.URL))
	if err != nil {
		t.Fatal(err)
	}
	var got bytes.Buffer
	if err := c.Timeseries.Download(context.Background(), timeseries.DataParams{Indicator: "x"}, &got); err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(got.Bytes(), data.Bytes()) {
		t.Fatal("ZIP changed")
	}
}
func TestDataValidation(t *testing.T) {
	c, err := wto.NewClient("key", wto.WithBaseURL("http://127.0.0.1:1"))
	if err != nil {
		t.Fatal(err)
	}
	for _, p := range []timeseries.DataParams{{}, {Indicator: "x", Format: "xml"}, {Indicator: "x", Mode: "invalid"}, {Indicator: "x", Max: ptr(0)}, {Indicator: "x", Offset: ptr(-1)}, {Indicator: "x", Headings: "H"}} {
		if _, err := c.Timeseries.Data(context.Background(), p); err == nil {
			t.Fatal(p)
		}
	}
}

func TestRawAndMalformedData(t *testing.T) {
	for _, post := range []bool{false, true} {
		t.Run(map[bool]string{true: "POST", false: "GET"}[post], func(t *testing.T) {
			const payload = `{"Dataset":[],"Summary":[{"Count":0}]}`
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.Write([]byte(payload)) }))
			defer server.Close()
			c, err := wto.NewClient("key", wto.WithBaseURL(server.URL))
			if err != nil {
				t.Fatal(err)
			}
			meta := true
			fn := c.Timeseries.DataRaw
			if post {
				fn = c.Timeseries.DataPostRaw
			}
			got, err := fn(context.Background(), timeseries.DataParams{Indicator: "x", Metadata: &meta})
			if err != nil || string(got) != payload {
				t.Fatal(string(got), err)
			}
		})
	}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.Write([]byte(`invalid`)) }))
	defer server.Close()
	c, err := wto.NewClient("key", wto.WithBaseURL(server.URL))
	if err != nil {
		t.Fatal(err)
	}
	if _, err := c.Timeseries.Data(context.Background(), timeseries.DataParams{Indicator: "x"}); err == nil {
		t.Fatal("malformed JSON accepted")
	}
}

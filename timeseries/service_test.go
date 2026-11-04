package timeseries_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	wto "github.com/cutedrag0n/wto-go"
	"github.com/cutedrag0n/wto-go/timeseries"
)

func ptr(v int) *int { return &v }
func TestOfficialOperations(t *testing.T) {
	t.Run("Topics", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/timeseries/v1/topics" || r.Method != "GET" {
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
		if _, err := c.Timeseries.Topics(context.Background(), timeseries.TopicsParams{}); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("Frequencies", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/timeseries/v1/frequencies" || r.Method != "GET" {
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
		if _, err := c.Timeseries.Frequencies(context.Background(), timeseries.FrequenciesParams{}); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("Periods", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/timeseries/v1/periods" || r.Method != "GET" {
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
		if _, err := c.Timeseries.Periods(context.Background(), timeseries.PeriodsParams{}); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("Units", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/timeseries/v1/units" || r.Method != "GET" {
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
		if _, err := c.Timeseries.Units(context.Background(), timeseries.UnitsParams{}); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("IndicatorCategories", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/timeseries/v1/indicator_categories" || r.Method != "GET" {
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
		if _, err := c.Timeseries.IndicatorCategories(context.Background(), timeseries.IndicatorCategoriesParams{}); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("Indicators", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/timeseries/v1/indicators" || r.Method != "GET" {
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
		if _, err := c.Timeseries.Indicators(context.Background(), timeseries.IndicatorsParams{}); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("Regions", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/timeseries/v1/territory/regions" || r.Method != "GET" {
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
		if _, err := c.Timeseries.Regions(context.Background(), timeseries.RegionsParams{}); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("Groups", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/timeseries/v1/territory/groups" || r.Method != "GET" {
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
		if _, err := c.Timeseries.Groups(context.Background(), timeseries.GroupsParams{}); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("Reporters", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/timeseries/v1/reporters" || r.Method != "GET" {
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
		if _, err := c.Timeseries.Reporters(context.Background(), timeseries.ReportersParams{}); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("Partners", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/timeseries/v1/partners" || r.Method != "GET" {
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
		if _, err := c.Timeseries.Partners(context.Background(), timeseries.PartnersParams{}); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("Classifications", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/timeseries/v1/product_classifications" || r.Method != "GET" {
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
		if _, err := c.Timeseries.Classifications(context.Background(), timeseries.ClassificationsParams{}); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("Products", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/timeseries/v1/products" || r.Method != "GET" {
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
		if _, err := c.Timeseries.Products(context.Background(), timeseries.ProductsParams{}); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("Years", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/timeseries/v1/years" || r.Method != "GET" {
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
		if _, err := c.Timeseries.Years(context.Background(), timeseries.YearsParams{}); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("ValueFlags", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/timeseries/v1/value_flags" || r.Method != "GET" {
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
		if _, err := c.Timeseries.ValueFlags(context.Background(), timeseries.ValueFlagsParams{}); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("Metadata", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/timeseries/v1/metadata" || r.Method != "GET" {
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
		if _, err := c.Timeseries.Metadata(context.Background(), timeseries.MetadataParams{Indicators: "ITS_MTV_AX"}); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("DataCount", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/timeseries/v1/data_count" || r.Method != "GET" {
				t.Errorf("unexpected request %s %s", r.Method, r.URL.Path)
			}
			if r.Header.Get("Ocp-Apim-Subscription-Key") != "test-key" {
				t.Error("missing key")
			}
			w.Write([]byte(`42`))
		}))
		defer server.Close()
		c, err := wto.NewClient("test-key", wto.WithBaseURL(server.URL))
		if err != nil {
			t.Fatal(err)
		}
		if _, err := c.Timeseries.DataCount(context.Background(), timeseries.DataCountParams{Indicators: "ITS_MTV_AX"}); err != nil {
			t.Fatal(err)
		}
	})
}

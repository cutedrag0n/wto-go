//go:build integration

package integration_test

import (
	"bytes"
	"context"
	"os"
	"testing"
	"time"

	wto "github.com/cutedrag0n/wto-go"
	"github.com/cutedrag0n/wto-go/eping"
	"github.com/cutedrag0n/wto-go/qr"
	"github.com/cutedrag0n/wto-go/tfad"
	"github.com/cutedrag0n/wto-go/timeseries"
)

func ptr(v int) *int { return &v }
func TestLive(t *testing.T) {
	key := os.Getenv("WTO_API_KEY")
	if key == "" {
		t.Skip("WTO_API_KEY is not set")
	}
	c, err := wto.NewClient(key, wto.WithTimeout(60*time.Second))
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	cases := []struct {
		name string
		run  func() error
	}{
		{"timeseries/Topics", func() error { _, err := c.Timeseries.Topics(ctx, timeseries.TopicsParams{}); return err }},
		{"timeseries/Frequencies", func() error { _, err := c.Timeseries.Frequencies(ctx, timeseries.FrequenciesParams{}); return err }},
		{"timeseries/Periods", func() error { _, err := c.Timeseries.Periods(ctx, timeseries.PeriodsParams{}); return err }},
		{"timeseries/Units", func() error { _, err := c.Timeseries.Units(ctx, timeseries.UnitsParams{}); return err }},
		{"timeseries/IndicatorCategories", func() error {
			_, err := c.Timeseries.IndicatorCategories(ctx, timeseries.IndicatorCategoriesParams{})
			return err
		}},
		{"timeseries/Indicators", func() error { _, err := c.Timeseries.Indicators(ctx, timeseries.IndicatorsParams{}); return err }},
		{"timeseries/Regions", func() error { _, err := c.Timeseries.Regions(ctx, timeseries.RegionsParams{}); return err }},
		{"timeseries/Groups", func() error { _, err := c.Timeseries.Groups(ctx, timeseries.GroupsParams{}); return err }},
		{"timeseries/Reporters", func() error { _, err := c.Timeseries.Reporters(ctx, timeseries.ReportersParams{}); return err }},
		{"timeseries/Partners", func() error { _, err := c.Timeseries.Partners(ctx, timeseries.PartnersParams{}); return err }},
		{"timeseries/Classifications", func() error {
			_, err := c.Timeseries.Classifications(ctx, timeseries.ClassificationsParams{})
			return err
		}},
		{"timeseries/Products", func() error { _, err := c.Timeseries.Products(ctx, timeseries.ProductsParams{}); return err }},
		{"timeseries/Years", func() error { _, err := c.Timeseries.Years(ctx, timeseries.YearsParams{}); return err }},
		{"timeseries/ValueFlags", func() error { _, err := c.Timeseries.ValueFlags(ctx, timeseries.ValueFlagsParams{}); return err }},
		{"timeseries/Metadata", func() error {
			_, err := c.Timeseries.Metadata(ctx, timeseries.MetadataParams{Indicators: "TP_A_0010", Reporters: "840"})
			return err
		}},
		{"timeseries/DataCount", func() error {
			_, err := c.Timeseries.DataCount(ctx, timeseries.DataCountParams{Indicators: "ITS_MTV_AX"})
			return err
		}},
		{"eping/SearchNotifications", func() error {
			_, err := c.EPing.SearchNotifications(ctx, eping.SearchNotificationsParams{Language: "1"})
			return err
		}},
		{"eping/Members", func() error { _, err := c.EPing.Members(ctx, eping.MembersParams{Language: "1"}); return err }},
		{"qr/Members", func() error { _, err := c.QR.Members(ctx, qr.MembersParams{}); return err }},
		{"qr/Products", func() error { _, err := c.QR.Products(ctx, qr.ProductsParams{HSVersion: "h6"}); return err }},
		{"qr/Notifications", func() error { _, err := c.QR.Notifications(ctx, qr.NotificationsParams{}); return err }},
		{"qr/Restriction", func() error { _, err := c.QR.Restriction(ctx, qr.RestrictionParams{QRID: ptr(14)}); return err }},
		{"qr/HSVersions", func() error { _, err := c.QR.HSVersions(ctx, qr.HSVersionsParams{}); return err }},
		{"qr/Restrictions", func() error { _, err := c.QR.Restrictions(ctx, qr.RestrictionsParams{}); return err }},
		{"tfad/Transparency", func() error { _, err := c.TFAD.Transparency(ctx, tfad.TransparencyParams{}); return err }},
		{"timeseries/Data", func() error {
			v, err := c.Timeseries.Data(ctx, timeseries.DataParams{Indicator: "ITS_MTV_AX", Max: ptr(1)})
			if err == nil && len(v) == 0 {
				t.Error("empty Dataset")
			}
			return err
		}},
		{"timeseries/DataPost", func() error {
			_, err := c.Timeseries.DataPost(ctx, timeseries.DataParams{Indicator: "ITS_MTV_AX", Max: ptr(1)})
			return err
		}},
		{"timeseries/Codes", func() error {
			_, err := c.Timeseries.Data(ctx, timeseries.DataParams{Indicator: "ITS_MTV_AX", Max: ptr(1), Mode: "codes"})
			return err
		}},
		{"timeseries/Download", func() error {
			var b bytes.Buffer
			err := c.Timeseries.Download(ctx, timeseries.DataParams{Indicator: "ITS_MTV_AX", Max: ptr(1)}, &b)
			if err == nil && !bytes.HasPrefix(b.Bytes(), []byte("PK")) {
				t.Error("not a ZIP")
			}
			return err
		}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if err := tc.run(); err != nil {
				t.Fatal(err)
			}
		})
	}
}

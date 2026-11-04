package timeseries

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/url"

	"github.com/cutedrag0n/wto-go/internal/transport"
)

// DataParams configures GET and POST data queries. Empty optional strings use server defaults.
// Use DataRaw or DataPostRaw for human headings, ZIP output, or metadata envelopes.
type DataParams struct {
	Indicator          string `json:"i"`
	Reporters          string `json:"r,omitempty"`
	Partners           string `json:"p,omitempty"`
	Periods            string `json:"ps,omitempty"`
	Products           string `json:"pc,omitempty"`
	IncludeSubproducts *bool  `json:"spc,omitempty"`
	Format             string `json:"fmt,omitempty"`
	Mode               string `json:"mode,omitempty"`
	Decimals           string `json:"dec,omitempty"`
	Offset             *int   `json:"off,omitempty"`
	Max                *int   `json:"max,omitempty"`
	Headings           string `json:"head,omitempty"`
	Language           *int   `json:"lang,omitempty"`
	Metadata           *bool  `json:"meta,omitempty"`
}

func (p DataParams) validate() error {
	if p.Indicator == "" {
		return fmt.Errorf("wto: indicator is required")
	}
	if p.Max != nil && (*p.Max < 1 || *p.Max > 1000000) {
		return fmt.Errorf("wto: max must be between 1 and 1000000")
	}
	if p.Offset != nil && (*p.Offset < 0 || *p.Offset > 1000000) {
		return fmt.Errorf("wto: offset must be between 0 and 1000000")
	}
	if p.Format != "" && p.Format != "json" && p.Format != "csv" {
		return fmt.Errorf("wto: format must be json or csv")
	}
	if p.Mode != "" && p.Mode != "full" && p.Mode != "codes" {
		return fmt.Errorf("wto: mode must be full or codes")
	}
	if p.Headings != "" && p.Headings != "M" && p.Headings != "H" {
		return fmt.Errorf("wto: headings must be M or H")
	}
	if p.Language != nil && (*p.Language < 1 || *p.Language > 3) {
		return fmt.Errorf("wto: language must be 1, 2 or 3")
	}
	return nil
}

func (p DataParams) query() url.Values {
	q := url.Values{}
	transport.Add(q, "i", p.Indicator)
	transport.Add(q, "r", p.Reporters)
	transport.Add(q, "p", p.Partners)
	transport.Add(q, "ps", p.Periods)
	transport.Add(q, "pc", p.Products)
	transport.Add(q, "spc", p.IncludeSubproducts)
	transport.Add(q, "fmt", p.Format)
	transport.Add(q, "mode", p.Mode)
	transport.Add(q, "dec", p.Decimals)
	transport.Add(q, "off", p.Offset)
	transport.Add(q, "max", p.Max)
	transport.Add(q, "head", p.Headings)
	transport.Add(q, "lang", p.Language)
	transport.Add(q, "meta", p.Metadata)
	return q
}

// Data returns machine-headed JSON datapoints. It accepts full and codes modes.
func (s *Service) Data(ctx context.Context, p DataParams) ([]DataPointExtended, error) {
	return s.data(ctx, p, false)
}

// DataPost performs the equivalent query with a JSON body to avoid URL length limits.
func (s *Service) DataPost(ctx context.Context, p DataParams) ([]DataPointExtended, error) {
	return s.data(ctx, p, true)
}
func (s *Service) data(ctx context.Context, p DataParams, post bool) ([]DataPointExtended, error) {
	if p.Format == "csv" || (p.Metadata != nil && *p.Metadata) || p.Headings == "H" {
		return nil, fmt.Errorf("wto: use raw data methods for CSV, metadata or human headings")
	}
	p.Format = "json"
	p.Headings = "M"
	b, err := s.raw(ctx, p, post)
	if err != nil {
		return nil, err
	}
	var response struct {
		Dataset []DataPointExtended `json:"Dataset"`
	}
	if err := json.Unmarshal(b, &response); err != nil {
		return nil, fmt.Errorf("wto: decode datapoints: %w", err)
	}
	return response.Dataset, nil
}

// DataRaw preserves any documented GET data output, including ZIP bytes.
func (s *Service) DataRaw(ctx context.Context, p DataParams) ([]byte, error) {
	return s.raw(ctx, p, false)
}

// DataPostRaw preserves any documented POST data output.
func (s *Service) DataPostRaw(ctx context.Context, p DataParams) ([]byte, error) {
	return s.raw(ctx, p, true)
}
func (s *Service) raw(ctx context.Context, p DataParams, post bool) ([]byte, error) {
	if err := p.validate(); err != nil {
		return nil, err
	}
	in := transport.Request{Service: "timeseries", Path: "/data", Method: "GET", Data: true}
	if post {
		in.Method = "POST"
		in.Body = p
	} else {
		in.Query = p.query()
	}
	return s.client.Do(ctx, in)
}

// Download writes the ZIP returned by fmt=csv. The response is buffered before writing.
func (s *Service) Download(ctx context.Context, p DataParams, w io.Writer) error {
	return s.download(ctx, p, w, false)
}

// DownloadPost downloads ZIP data using the POST operation.
func (s *Service) DownloadPost(ctx context.Context, p DataParams, w io.Writer) error {
	return s.download(ctx, p, w, true)
}
func (s *Service) download(ctx context.Context, p DataParams, w io.Writer, post bool) error {
	if w == nil {
		return fmt.Errorf("wto: writer is required")
	}
	p.Format = "csv"
	b, err := s.raw(ctx, p, post)
	if err != nil {
		return err
	}
	n, err := w.Write(b)
	if err != nil {
		return fmt.Errorf("wto: write ZIP: %w", err)
	}
	if n != len(b) {
		return fmt.Errorf("wto: write ZIP: %w", io.ErrShortWrite)
	}
	return nil
}

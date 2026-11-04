# wto-go

An unofficial, typed Go SDK for the [World Trade Organization APIs](https://apiportal.wto.org/apis), built with Resty v2. Requires Go 1.24 or newer. [中文文档](README.zh-CN.md).

## Install and authenticate

```sh
go get github.com/cutedrag0n/wto-go
export WTO_API_KEY='your-subscription-key'
```

Register at the [WTO developer portal](https://apiportal.wto.org/signup), subscribe to the free product, and obtain your subscription key. The SDK sends it in `Ocp-Apim-Subscription-Key`; it does not read environment variables automatically.

```go
import (
    "context"
    "os"
    "time"

    wto "github.com/cutedrag0n/wto-go"
    "github.com/cutedrag0n/wto-go/timeseries"
)

client, err := wto.NewClient(os.Getenv("WTO_API_KEY"), wto.WithTimeout(time.Minute))
if err != nil { /* handle error */ }
limit := 100
rows, err := client.Timeseries.Data(context.Background(), timeseries.DataParams{
    Indicator: "ITS_MTV_AX",
    Reporters: "840",
    Periods: "2020-2023",
    Max: &limit,
})
```

`Data` and `DataPost` return datapoints from the actual `Dataset` envelope, force `fmt=json` and `head=M`, and accept `Mode: "full"` or `"codes"`. `DataPost` sends the same filters in a JSON body for long queries. Nullable data fields use pointers. To paginate, set `Offset` and `Max` explicitly; the SDK never automatically fetches all records.

## Services

Import parameter types from the service package you use:

```go
// Imports: .../eping, .../qr, .../tfad
notifications, err := client.EPing.SearchNotifications(ctx, eping.SearchNotificationsParams{
    Language: "1", CountryIDs: []string{"C840"},
})
products, err := client.QR.Products(ctx, qr.ProductsParams{HSVersion: "h6", Code: "0101"})
transparency, err := client.TFAD.Transparency(ctx, tfad.TransparencyParams{
    Countries: []string{"004", "008"},
})
```

All 27 official operations are listed in [API coverage](docs/api-coverage.md). QR lists return typed `Page[T]` values including `Data`, `Links`, and `Meta`. Follow pages through the documented page parameter, rather than sending credentials to URLs returned in pagination links.

## Raw output and downloads

```go
includeMetadata := true
raw, err := client.Timeseries.DataRaw(ctx, timeseries.DataParams{
    Indicator: "ITS_MTV_AX", Metadata: &includeMetadata,
})
// raw preserves the complete JSON output including Summary and metadata when present.

err = client.Timeseries.Download(ctx, timeseries.DataParams{
    Indicator: "ITS_MTV_AX", Max: &limit,
}, outputFile) // io.Writer; fmt=csv returns ZIP bytes
```

`DataPostRaw` and `DownloadPost` provide POST equivalents. Downloads buffer the response before writing; limit the query size for large datasets. Human headings (`head=H`), metadata and ZIP formats are available through raw methods. `Data`/`DataPost` reject these incompatible output options instead of silently discarding them.

## Configuration and rate limits

```go
limiter := wto.NewRateLimiter()
client, err := wto.NewClient(key,
    wto.WithRateLimiter(limiter),
    wto.WithTimeout(2*time.Minute),
    wto.WithRetry(2), // additional attempts, default 0
    wto.WithProxy("http://localhost:8080"),
    wto.WithHeader("User-Agent", "my-app"),
)
// Reuse limiter for every Client consuming the same subscription in this process.
```

Rate limiting is enabled by default:

| Limit | Local enforcement |
|---|---|
| 10 calls/second | Token bucket, burst 1; smooth admission |
| 10,000 calls/hour | Rolling one-hour quota |
| Timeseries data + data_count: 1 call/second | Shared token bucket, burst 1, including POST |

Every attempt, including retries, consumes quota. Waiting is canceled by Context. The default **total call budget is 30 seconds**, including quota waiting, HTTP work and retry delays. Increase it or provide an earlier Context deadline as needed. Hourly exhaustion waits until quota becomes available or the call times out.

429 responses with a valid `Retry-After` establish a cooldown shared by clients using that limiter. Both delta seconds and HTTP dates are supported. Optional retries cover temporary network failures and HTTP 429/502/503/504, with bounded exponential backoff. Every retry passes through the limiter again.

Quota state is in memory: restarting, separate processes or clients without a shared limiter cannot account for each other's calls. The hourly rolling window is a conservative default because the published rule does not specify a reset boundary. A free subscription still has these limits.

`WithHTTPClient` copies `http.Client` settings and shares its Transport and Jar. Proxy configuration clones a standard `*http.Transport`; it rejects incompatible custom RoundTrippers. Cross-origin redirects are stopped to protect subscription headers; same-origin redirects honor the supplied client's redirect callback. `WithBaseURL` overrides the gateway while retaining service prefixes; `WithServiceBaseURL(wto.QR, url)` overrides a complete service base URL.

## Errors and documented gaps

```go
var apiErr *wto.APIError
if errors.As(err, &apiErr) {
    // Inspect StatusCode, Errors, Title, Detail, Message and Body.
}
if errors.Is(err, context.DeadlineExceeded) {
    // Total call budget or caller deadline expired.
}
```

Error strings omit server content to avoid reflecting credentials. `APIError.Body` preserves non-JSON errors too. Transport errors support unwrapping. HTTP error codes are preserved, including validation 400/422, authentication 401/403, not found 404 and throttling 429.

The official ePing Members and TFAD definitions do not describe response models: these return `json.RawMessage`. Unspecified nested fields also remain raw JSON. [API notes](docs/api-notes.md) records schema conflicts and live observations; the SDK does not claim full typing where the official contract is incomplete.

## Tests

```sh
go test ./...
go test -race -cover ./...
go vet ./...
go build ./...
WTO_API_KEY='your-key' go test -tags=integration -v ./integration
```

See the [validation report](docs/testing.md) for tested scenarios and live results. Unit tests are offline. Live tests skip when the key is absent and consume subscription quota when enabled. Examples are in [examples](examples/). Official definition snapshots and provenance are in [docs/openapi](docs/openapi/).

MIT licensed. This project is not affiliated with or endorsed by the WTO.

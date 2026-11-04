# Validation report

Validated on 2026-11-04 using Go 1.24.6 (darwin/arm64).

| Check | Result |
|---|---|
| Offline unit tests with race detector | PASS |
| Aggregate SDK statement coverage | 93.0% |
| go vet ./... | PASS |
| go build ./... (including all examples) | PASS |
| git diff --check | PASS |
| Live integration | 29 scenarios PASS, covering all 27 official operations plus codes/ZIP variants |
| Project scan for supplied test credential | 0 occurrences |

Coverage was collected across the SDK packages (root, timeseries, eping, qr, tfad, internal/transport), excluding example main packages:

```sh
go test -race \
  -coverpkg=.,./timeseries,./eping,./qr,./tfad,./internal/transport \
  -coverprofile=/tmp/wto-coverage.out ./...
go tool cover -func=/tmp/wto-coverage.out
go vet ./...
go build ./...
WTO_API_KEY='your-key' go test -tags=integration -count=1 -v ./integration
```

Offline coverage includes operation paths/methods/authentication; explicit false and zero values; repeated array query keys; POST JSON; Dataset decoding in full/codes modes; nullable fields; original output preservation; ZIP integrity; official QR and ePing response fixtures; numeric/string scalar compatibility; validation; HTTP error payloads; malformed/non-JSON output; deadlines; retry attempts and cancellation; Retry-After formats; shared 429 cooldown; custom HTTP transport; proxy routing; redirect credential isolation; and shared client limiter behavior.

Rate limiter boundary tests use explicit timestamps passed to the admission calculation: combined global/data limits, 10,000-entry hourly exhaustion, exact expiry, cancellation without quota consumption, and concurrent admission. They do not wait for one hour or intentionally overload the live API. Same-origin redirects are also counted and paced.

Live tests read WTO_API_KEY at runtime, run sequentially through the production limiter, and skip without a key. They establish success for the tested inputs, not all possible parameter combinations. An initial metadata call using ITS_MTV_AX returned server-side 500; the official default indicator TP_A_0010 with reporter 840 succeeded and is used in the regression smoke test. See api-notes.md for other upstream inconsistencies.

Credentials are not saved in code, fixtures, snapshots, or this report.

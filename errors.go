package wto

import "github.com/cutedrag0n/wto-go/internal/transport"

// APIError describes a non-success HTTP response. Use errors.As to inspect it.
type APIError = transport.APIError

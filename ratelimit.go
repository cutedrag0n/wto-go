package wto

import "github.com/cutedrag0n/wto-go/internal/transport"

// RateLimiter coordinates a subscription's request attempts within one process.
type RateLimiter = transport.RateLimiter

// NewRateLimiter returns a limiter enforcing all three WTO subscription limits.
func NewRateLimiter() *RateLimiter { return transport.NewRateLimiter() }

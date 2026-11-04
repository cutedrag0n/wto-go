package transport

import (
	"context"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

// RateLimiter coordinates all attempts using one WTO subscription. Its zero value is ready to use.
type RateLimiter struct {
	mu           sync.Mutex
	global, data *rate.Limiter
	sent         []time.Time
	cooldown     time.Time
}

func NewRateLimiter() *RateLimiter { return &RateLimiter{} }

// delayLocked atomically admits an attempt only when every limit permits it.
func (l *RateLimiter) delayLocked(now time.Time, data bool) time.Duration {
	if l.global == nil {
		l.global = rate.NewLimiter(10, 1)
		l.data = rate.NewLimiter(1, 1)
	}
	expired := 0
	for expired < len(l.sent) && !l.sent[expired].Add(time.Hour).After(now) {
		expired++
	}
	l.sent = l.sent[expired:]
	delay := max(time.Duration(0), l.cooldown.Sub(now))
	if len(l.sent) >= 10000 {
		delay = max(delay, l.sent[0].Add(time.Hour).Sub(now))
	}
	delay = max(delay, tokenDelay(l.global, now))
	if data {
		delay = max(delay, tokenDelay(l.data, now))
	}
	if delay > 0 {
		return delay
	}
	// All consumers access these buckets under mu; no reservation can race admission.
	l.global.AllowN(now, 1)
	if data {
		l.data.AllowN(now, 1)
	}
	l.sent = append(l.sent, now)
	return 0
}

func tokenDelay(l *rate.Limiter, now time.Time) time.Duration {
	missing := 1 - l.TokensAt(now)
	if missing <= 0 {
		return 0
	}
	return time.Duration(missing/float64(l.Limit())*float64(time.Second)) + time.Nanosecond
}

func (l *RateLimiter) Wait(ctx context.Context, data bool) error {
	for {
		if err := ctx.Err(); err != nil {
			return err
		}
		l.mu.Lock()
		if err := ctx.Err(); err != nil {
			l.mu.Unlock()
			return err
		}
		delay := l.delayLocked(time.Now(), data)
		l.mu.Unlock()
		if delay == 0 {
			return nil
		}
		if err := sleep(ctx, delay); err != nil {
			return err
		}
	}
}

func (l *RateLimiter) Cooldown(until time.Time) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if until.After(l.cooldown) {
		l.cooldown = until
	}
}

func sleep(ctx context.Context, d time.Duration) error {
	timer := time.NewTimer(d)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}

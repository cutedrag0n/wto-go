package transport

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"
)

func TestAdmission(t *testing.T) {
	now := time.Now()
	l := NewRateLimiter()
	l.mu.Lock()
	defer l.mu.Unlock()
	if d := l.delayLocked(now, true); d != 0 {
		t.Fatal(d)
	}
	if d := l.delayLocked(now.Add(100*time.Millisecond), true); d < 899*time.Millisecond {
		t.Fatal("data spacing", d)
	}
	if d := l.delayLocked(now.Add(100*time.Millisecond), false); d != 0 {
		t.Fatal("ordinary request blocked", d)
	}
	if len(l.sent) != 2 {
		t.Fatal(l.sent)
	}
	if d := l.delayLocked(now.Add(time.Second), true); d != 0 {
		t.Fatal(d)
	}
}
func TestHourlyQuota(t *testing.T) {
	l := NewRateLimiter()
	now := time.Now()
	l.sent = make([]time.Time, 10000)
	for i := range l.sent {
		l.sent[i] = now.Add(-time.Minute)
	}
	if d := l.delayLocked(now, false); d != 59*time.Minute {
		t.Fatal(d)
	}
	if len(l.sent) != 10000 {
		t.Fatal("waiting consumed quota")
	}
	if d := l.delayLocked(now.Add(59*time.Minute), false); d != 0 {
		t.Fatal(d)
	}
	if len(l.sent) != 1 {
		t.Fatal(len(l.sent))
	}
}
func TestWaitCancellationAndCooldown(t *testing.T) {
	l := NewRateLimiter()
	l.Cooldown(time.Now().Add(time.Hour))
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	defer cancel()
	if err := l.Wait(ctx, false); !errors.Is(err, context.DeadlineExceeded) {
		t.Fatal(err)
	}
	if len(l.sent) != 0 {
		t.Fatal("cancellation consumed quota")
	}
}
func TestConcurrentAdmission(t *testing.T) {
	l := NewRateLimiter()
	now := time.Now()
	var wg sync.WaitGroup
	admitted := 0
	for range 100 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			l.mu.Lock()
			defer l.mu.Unlock()
			if l.delayLocked(now, false) == 0 {
				admitted++
			}
		}()
	}
	wg.Wait()
	if admitted != 1 {
		t.Fatal(admitted)
	}
}

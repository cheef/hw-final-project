package services

import (
	"log/slog"
	"math"
	"sync"
	"time"
)

type TokenBucket struct {
	tokens    float64
	MaxTokens int
	Duration  int
	done      chan struct{}
	mu        sync.Mutex
	LastUsage time.Time
	log       *slog.Logger
}

func NewTokenBucket(maxTokens, duration int, log *slog.Logger) *TokenBucket {
	bucket := TokenBucket{
		tokens:    float64(maxTokens),
		MaxTokens: maxTokens,
		Duration:  duration,
		LastUsage: time.Now(),
		done:      make(chan struct{}),
		log:       log,
	}

	go bucket.refillTokens()

	return &bucket
}

func (b *TokenBucket) refillTokens() {
	tickerTime := time.Duration(b.Duration / b.MaxTokens)
	ticker := time.NewTicker(tickerTime * time.Millisecond)

	b.log.Debug(
		"TokenBucket refill tokens",
		slog.Int("Duration", b.Duration),
		slog.Int("MaxTokens", b.MaxTokens),
		slog.Duration("tickerTime (ms)", tickerTime),
	)

	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			b.mu.Lock()
			if b.tokens != float64(b.MaxTokens) {
				b.tokens = math.Min(b.tokens+1, float64(b.MaxTokens))
				b.log.Debug(
					"TokenBucket added token",
					slog.Float64("left", b.tokens),
				)
			}
			b.mu.Unlock()
		case <-b.done:
			return
		}
	}
}

func (b *TokenBucket) Stop() {
	close(b.done)
}

func (b *TokenBucket) UseToken() bool {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.LastUsage = time.Now()

	if b.tokens < 1 {
		b.log.Debug(
			"TokenBucket no free tokens are left",
			slog.Float64("left", b.tokens),
		)
		return false
	}

	b.tokens--
	b.log.Debug(
		"TokenBucket uses token",
		slog.Float64("left", b.tokens),
	)

	return true
}

func (b *TokenBucket) GetLastAccess() time.Time {
	b.mu.Lock()
	defer b.mu.Unlock()

	return b.LastUsage
}

package services

import (
	"log/slog"
	"sync"
	"time"
)

type Credential string

type CredentialLimiter struct {
	buckets        map[Credential]*TokenBucket
	MaxTokens      int
	Duration       int
	BucketLifeTime int
	done           chan struct{}
	mu             sync.Mutex
	log            *slog.Logger
}

func NewCredentialLimiter(maxTokens, duration, bucketLifeTime int, log *slog.Logger) *CredentialLimiter {
	buckets := make(map[Credential]*TokenBucket)

	log.Debug(
		"CredentialLimiter built",
		slog.Int("maxTokens", maxTokens),
		slog.Int("duration", duration),
		slog.Int("bucketLifeTime", bucketLifeTime),
	)

	return &CredentialLimiter{
		MaxTokens:      maxTokens,
		Duration:       duration,
		BucketLifeTime: bucketLifeTime,
		buckets:        buckets,
		done:           make(chan struct{}),
		log:            log,
	}
}

func (ct *CredentialLimiter) IsAllowed(c Credential) bool {
	bucket, ok := ct.GetBucket(c)

	if ok {
		return bucket.UseToken()
	}

	bucket = ct.addBucket(c)

	return bucket.UseToken()
}

func (ct *CredentialLimiter) buildBucket(c Credential) *TokenBucket {
	bucket := NewTokenBucket(ct.MaxTokens, ct.Duration, ct.log)

	ct.log.Debug(
		"CredentialLimiter creating new bucket",
		slog.String("key", string(c)),
		slog.Int("maxTokens", ct.MaxTokens),
		slog.Int("duration", ct.Duration),
	)

	return bucket
}

func (ct *CredentialLimiter) addBucket(c Credential) *TokenBucket {
	bucket := ct.buildBucket(c)

	ct.mu.Lock()
	ct.buckets[c] = bucket
	ct.mu.Unlock()

	if ct.BucketLifeTime > 0 {
		go ct.sweepBucket(c)
	}

	return bucket
}

func (ct *CredentialLimiter) GetBucket(c Credential) (*TokenBucket, bool) {
	ct.mu.Lock()
	bucket, ok := ct.buckets[c]
	ct.mu.Unlock()

	return bucket, ok
}

func (ct *CredentialLimiter) RemoveBucket(c Credential) {
	bucket, ok := ct.GetBucket(c)

	if !ok {
		return
	}

	bucket.Stop()

	ct.mu.Lock()
	delete(ct.buckets, c)
	ct.mu.Unlock()
}

func (ct *CredentialLimiter) sweepBucket(c Credential) {
	duration := time.Duration(ct.BucketLifeTime) * time.Millisecond
	ticker := time.NewTicker(duration)

	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			bucket, ok := ct.GetBucket(c)

			if !ok {
				return
			}

			sweepTime := bucket.GetLastAccess()

			if time.Now().After(sweepTime) {
				ct.log.Debug("CredentialLimiter removing stale bucket", slog.String("key", string(c)))
				ct.RemoveBucket(c)
				return
			}
		case <-ct.done:
			return
		}
	}
}

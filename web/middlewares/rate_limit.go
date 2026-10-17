package middlewares

import (
	"hash/fnv"
	"net/http"
	"sync"
	"time"

	"github.com/arthurcommodore/components/web/utils"
)

type visitor struct {
	tokens     int
	lastRefill time.Time
	lastSeen   time.Time
}

type bucket struct {
	mu       sync.RWMutex
	visitors map[string]*visitor
}

type rateLimiter struct {
	buckets   []bucket
	maxTokens int
	refill    time.Duration
	ttl       time.Duration
}

func NewRateLimit(maxTokens int, refill time.Duration, ttl time.Duration) *rateLimiter {
	return &rateLimiter{
		buckets:   make([]bucket, 32),
		maxTokens: maxTokens,
		refill:    refill,
		ttl:       ttl,
	}
}

func (rl *rateLimiter) getBucket(ip string) *bucket {
	h := fnv.New32a()
	h.Write([]byte(ip))
	idx := int(h.Sum32()) % len(rl.buckets)
	return &rl.buckets[idx]
}

func (rl *rateLimiter) Allow(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := utils.GetClientIP(r)
		if ip == "" {
			http.Error(w, "Unauthorized IP", http.StatusUnauthorized)
			return
		}

		b := rl.getBucket(ip)
		b.mu.Lock()
		defer b.mu.Unlock()

		if b.visitors == nil {
			b.visitors = make(map[string]*visitor)
		}

		v := b.visitors[ip]

		if v == nil {
			v = &visitor{
				tokens:     rl.maxTokens,
				lastRefill: time.Now(),
				lastSeen:   time.Now(),
			}
			b.visitors[ip] = v
		}

		if v.lastRefill.Add(rl.refill).Before(time.Now()) {
			v.lastRefill = time.Now()
			v.tokens = rl.maxTokens
		}

		v.lastSeen = time.Now()
		v.tokens -= 1
		if v.tokens < 0 {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}

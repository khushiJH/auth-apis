package utils

import "sync"

var (
	blacklistedTokens = make(map[string]bool)
	mu                sync.RWMutex
)

func BlacklistToken(token string) {
	mu.Lock()
	defer mu.Unlock()
	blacklistedTokens[token] = true
}

func IsTokenBlacklisted(token string) bool {
	mu.RLock()
	defer mu.RUnlock()
	return blacklistedTokens[token]
}

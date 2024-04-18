package limiter

import (
	"context"
	"sync"

	"golang.org/x/time/rate"
)

// Limiter 限流
type Limiter struct {
	apiLimit map[string]*rate.Limiter
	locker   sync.Mutex
}

// New *Limiter
func New() *Limiter {
	return &Limiter{
		apiLimit: map[string]*rate.Limiter{},
		locker:   sync.Mutex{},
	}
}

// SetNX 如果没有控制则添加控制，否则不做处理
func (lim *Limiter) SetNX(key string, l *rate.Limiter) {
	if _, ok := lim.apiLimit[key]; !ok {
		lim.locker.Lock()
		defer lim.locker.Unlock()
		lim.apiLimit[key] = l
	}
}

// Allow 不在控制范围内的都直接放过
func (lim *Limiter) Allow(key string) bool {
	if tmp, ok := lim.apiLimit[key]; ok {
		return tmp.Allow()
	}
	return true
}

// Wait 等待获取令牌
func (lim *Limiter) Wait(ctx context.Context, key string) error {
	if tmp, ok := lim.apiLimit[key]; ok {
		return tmp.Wait(ctx)
	}
	return nil
}

// PrintKey 等待获取令牌
func (lim *Limiter) PrintKey() []string {
	keys := make([]string, 0, len(lim.apiLimit))
	for k := range lim.apiLimit {
		keys = append(keys, k)
	}
	return keys
}

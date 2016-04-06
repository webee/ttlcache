package ttlcache

import (
	"sync"
	"time"
)

const (
	entryNotExpire           time.Duration = -1
	entryExpireWithGlobalTTL time.Duration = 0
)

func newEntry(key string, content interface{}, ttl time.Duration) *entry {
	e := &entry{
		content: content,
		ttl:     ttl,
		key:     key,
	}
	e.touch()
	return e
}

type entry struct {
	key        string
	content    interface{}
	ttl        time.Duration
	expireAt   time.Time
	mutex      sync.Mutex
	queueIndex int
}

func (e *entry) touch() {
	e.mutex.Lock()
	if e.ttl > 0 {
		e.expireAt = time.Now().Add(item.ttl)
	}
	e.mutex.Unlock()
}

func (e *entry) expired() bool {
	e.mutex.Lock()
	if e.ttl <= 0 {
		e.mutex.Unlock()
		return false
	}
	expired := e.expireAt.Before(time.Now())
	e.mutex.Unlock()
	return expired
}

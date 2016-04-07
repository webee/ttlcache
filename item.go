package ttlcache

import (
	"sync"
	"time"
)

func newEntry(key string, content interface{}, ttl time.Duration) *entry {
	e := &entry{
		key:     key,
		content: content,
		ttl:     ttl,
	}
	e.touch()
	return e
}

type entry struct {
	key      string
	content  interface{}
	ttl      time.Duration
	expireAt time.Time
	mutex    sync.Mutex
}

func (e *entry) touch() {
	e.mutex.Lock()
	if e.ttl > 0 {
		e.expireAt = time.Now().Add(e.ttl)
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

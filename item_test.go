package ttlcache

import (
	"testing"
	"time"
)

func TestEntryExpired(t *testing.T) {
	e := newEntry("key", "value", (time.Duration(100) * time.Millisecond))
	if e.expired() {
		t.Error("Expected item to not be expired")
	}

	<-time.After(200 * time.Millisecond)

	if !e.expired() {
		t.Error("Expected item to be expired")
	}
}

func TestEntryTouch(t *testing.T) {
	e := newEntry("key", "value", (time.Duration(100) * time.Millisecond))
	oldExpireAt := e.expireAt
	<-time.After(50 * time.Millisecond)
	e.touch()

	if oldExpireAt == e.expireAt {
		t.Error("Expected dates to be different")
	}

	<-time.After(150 * time.Millisecond)

	if !e.expired() {
		t.Error("Expected item to be expired")
	}

	e.touch()
	<-time.After(50 * time.Millisecond)

	if e.expired() {
		t.Error("Expected item to not be expired")
	}
}

func TestEntryWithoutExpiration(t *testing.T) {
	e := newEntry("key", "value", 0)
	<-time.After(50 * time.Millisecond)
	if e.expired() {
		t.Error("Expected item to not be expired")
	}
}

func BenchmarkEntryNew(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		newEntry("key", "value", -1)
	}
}

func BenchmarkEntryTouch(b *testing.B) {
	b.ReportAllocs()
	e := newEntry("key", "value", -1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		e.touch()
	}
}

func BenchmarkEntryExpired(b *testing.B) {
	b.ReportAllocs()
	e := newEntry("key", "value", -1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		e.expired()
	}
}

package ks

import (
	"sync"
	"time"
)

const (
	defaultInitialHashMapUnderlyingArraySize int = 16
	defaultListCapacity                      int = 100
)

type KS struct {
	hashMap *hashMap
	list    *list
	mu      sync.Mutex
}

func newKS() *KS {
	return &KS{
		hashMap: newHashMap(FowlerNollVoHashFunction, defaultInitialHashMapUnderlyingArraySize),
		list:    newList(defaultListCapacity),
	}
}

func (ks *KS) Set(key string, value string, ttl int) {
	ks.mu.Lock()
	defer ks.mu.Unlock()

	// cheapest way to find out if the key already exists and
	// we need to update the value and move the key to the head of the list
	// instead of inserting it at the head of the list
	if _, ok := ks.hashMap.get(key); ok {
		ks.list.delete(key)
	}

	if removedKey := ks.list.insertHead(key); removedKey != "" {
		ks.hashMap.delete(removedKey)
	}

	ks.hashMap.set(key, value, time.Duration(ttl)*time.Second)
}

func (ks *KS) Get(key string) (string, bool) {
	ks.mu.Lock()
	defer ks.mu.Unlock()
	return ks.hashMap.get(key)
}

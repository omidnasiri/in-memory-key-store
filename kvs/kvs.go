package kvs

import (
	"sync"
	"time"
)

const (
	defaultInitialHashMapUnderlyingArraySize int = 16
	defaultListCapacity                      int = 100
)

type KVS struct {
	hashMap *hashMap
	list    *list
	mu      sync.Mutex
}

func newKVS() *KVS {
	return &KVS{
		hashMap: newHashMap(FowlerNollVoHashFunction, defaultInitialHashMapUnderlyingArraySize),
		list:    newList(defaultListCapacity),
	}
}

func (kvs *KVS) Set(key string, value string, ttl int) {
	kvs.mu.Lock()
	defer kvs.mu.Unlock()

	// cheapest way to find out if the key already exists and
	// we need to update the value and move the key to the head of the list
	// instead of inserting it at the head of the list
	if _, ok := kvs.hashMap.get(key); ok {
		kvs.list.delete(key)
	}

	if removedKey := kvs.list.insertHead(key); removedKey != "" {
		kvs.hashMap.delete(removedKey)
	}

	kvs.hashMap.set(key, value, time.Duration(ttl)*time.Second)
}

func (kvs *KVS) Get(key string) (string, bool) {
	kvs.mu.Lock()
	defer kvs.mu.Unlock()
	return kvs.hashMap.get(key)
}

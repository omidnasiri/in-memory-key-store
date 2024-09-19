package main

import "time"

type hashMapNode struct {
	key    string
	value  string
	expiry time.Time
}

type HashMap struct {
	population int
	hashFunc   func(key string, limit int) int

	// Using 2D slice to implement chaining in response to the collision problem.
	bucket [][]*hashMapNode
}

func NewHashMap(hashFunc func(key string, limit int) int, initialUnderlyingArraySize int) *HashMap {
	return &HashMap{
		bucket:   make([][]*hashMapNode, initialUnderlyingArraySize),
		hashFunc: hashFunc,
	}
}

func (hm *HashMap) Set(key string, value string, ttl time.Duration) {
	index := hm.hashFunc(key, len(hm.bucket))

	node := &hashMapNode{
		key:    key,
		value:  value,
		expiry: time.Now().Add(ttl),
	}

	hm.bucket[index] = append(hm.bucket[index], node)
	hm.population++
}

func (hm *HashMap) Get(key string) (string, bool) {
	index := hm.hashFunc(key, len(hm.bucket))

	for _, node := range hm.bucket[index] {
		if node.key == key {
			if node.expiry.Before(time.Now()) {
				return "", false
			}
			return node.value, true
		}
	}

	return "", false
}

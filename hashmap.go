package main

import "time"

type hashMapNode struct {
	key    string
	value  string
	expiry time.Time
}

type HashMap struct {
	population int
	hashFunc   HashFunction

	// Using 2D slice to implement chaining in response to the collision problem.
	bucket [][]*hashMapNode
}

func NewHashMap(hashFunc HashFunction, initialUnderlyingArraySize int) *HashMap {
	return &HashMap{
		bucket:   make([][]*hashMapNode, initialUnderlyingArraySize),
		hashFunc: hashFunc,
	}
}

func (hm *HashMap) Set(key string, value string, ttl time.Duration) {
	index := hm.hashFunc(key, len(hm.bucket))

	for _, node := range hm.bucket[index] {
		if node.key == key {
			node.value = value
			node.expiry = time.Now().Add(ttl)
			return
		}
	}

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
				// could also remove the node here
				return "", false
			}
			return node.value, true
		}
	}

	return "", false
}

func (hm *HashMap) Delete(key string) {
	index := hm.hashFunc(key, len(hm.bucket))

	for i, node := range hm.bucket[index] {
		if node.key == key {
			hm.bucket[index] = append(hm.bucket[index][:i], hm.bucket[index][i+1:]...)
			hm.population--
			return
		}
	}
}

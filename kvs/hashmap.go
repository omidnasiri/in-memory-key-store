package kvs

import "time"

type hashMapNode struct {
	key    string
	value  string
	expiry time.Time
}

type hashMap struct {
	population int
	hashFunc   HashFunction

	// Using 2D slice to implement chaining in response to the collision problem.
	bucket [][]*hashMapNode
}

func newHashMap(hashFunc HashFunction, initialUnderlyingArraySize int) *hashMap {
	return &hashMap{
		bucket:   make([][]*hashMapNode, initialUnderlyingArraySize),
		hashFunc: hashFunc,
	}
}

func (hm *hashMap) set(key string, value string, ttl time.Duration) {
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

func (hm *hashMap) get(key string) (string, bool) {
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

func (hm *hashMap) delete(key string) {
	index := hm.hashFunc(key, len(hm.bucket))

	for i, node := range hm.bucket[index] {
		if node.key == key {
			hm.bucket[index] = append(hm.bucket[index][:i], hm.bucket[index][i+1:]...)
			hm.population--
			return
		}
	}
}

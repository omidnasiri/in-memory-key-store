package ks

import "time"

const (
	defaultInitialHashMapUnderlyingArraySize int = 16
	defaultListCapacity                      int = 100
)

type KS struct {
	HashMap *hashMap
	List    *list
}

func NewKS() *KS {
	return &KS{
		HashMap: NewHashMap(FowlerNollVoHashFunction, defaultInitialHashMapUnderlyingArraySize),
		List:    NewList(defaultListCapacity),
	}
}

func (ks *KS) Set(key string, value string, ttl int) {
	// cheapest way to find out if the key already exists and
	// we need to update the value and move the key to the head of the list
	// instead of inserting it at the head of the list
	if _, ok := ks.HashMap.Get(key); ok {
		ks.List.Delete(key)
	}

	if removedKey := ks.List.InsertHead(key); removedKey != "" {
		ks.HashMap.Delete(removedKey)
	}

	ks.HashMap.Set(key, value, time.Duration(ttl)*time.Second)
}

func (ks *KS) Get(key string) (string, bool) {
	return ks.HashMap.Get(key)
}

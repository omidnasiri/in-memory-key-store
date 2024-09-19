package main

import "time"

const (
	defaultInitialHashMapUnderlyingArraySize int = 16
)

type KS struct {
	HashMap *HashMap
	List    any
}

func NewKS() *KS {
	return &KS{
		HashMap: NewHashMap(func(key string, limit int) int { return 0 }, defaultInitialHashMapUnderlyingArraySize),
	}
}

func (ks *KS) Set(key string, value string, ttl int) {
	ks.HashMap.Set(key, value, time.Duration(ttl)*time.Second)
}

func (ks *KS) Get(key string) (string, bool) {
	return ks.HashMap.Get(key)
}

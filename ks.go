package main

import "time"

const (
	defaultInitialHashMapUnderlyingArraySize int = 16
	defaultListCapacity                      int = 100
)

type KS struct {
	HashMap *HashMap
	List    *List
}

func NewKS() *KS {
	return &KS{
		HashMap: NewHashMap(FowlerNollVoHashFunction, defaultInitialHashMapUnderlyingArraySize),
		List:    NewList(defaultListCapacity),
	}
}

func (ks *KS) Set(key string, value string, ttl int) {
	ks.HashMap.Set(key, value, time.Duration(ttl)*time.Second)
}

func (ks *KS) Get(key string) (string, bool) {
	return ks.HashMap.Get(key)
}

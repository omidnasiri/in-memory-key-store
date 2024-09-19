package main

type KS struct {
	HashMap any
	List    any
}

func NewKS() *KS {
	return &KS{}
}

func (ks *KS) Set(key string, value string, ttl int) {}

func (ks *KS) Get(key string) (string, bool) {
	return "", false
}

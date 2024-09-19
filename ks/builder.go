package ks

type KSBuilder struct {
	ks *KS
}

func NewKSBuilder() *KSBuilder {
	return &KSBuilder{
		ks: newKS(),
	}
}

func (ksb *KSBuilder) WithHashFunction(hashFunc HashFunction) *KSBuilder {
	ksb.ks.hashMap.hashFunc = hashFunc
	return ksb
}

func (ksb *KSBuilder) WithInitialHashMapUnderlyingArraySize(size int) *KSBuilder {
	ksb.ks.hashMap.bucket = make([][]*hashMapNode, size)
	return ksb
}

func (ksb *KSBuilder) WithListCapacity(capacity int) *KSBuilder {
	ksb.ks.list.capacity = capacity
	return ksb
}

func (ksb *KSBuilder) Build() *KS {
	return ksb.ks
}

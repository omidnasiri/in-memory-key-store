package ks

type KSBuilder struct {
	ks *KS
}

func NewKSBuilder() *KSBuilder {
	return &KSBuilder{
		ks: NewKS(),
	}
}

func (ksb *KSBuilder) WithHashFunction(hashFunc HashFunction) *KSBuilder {
	ksb.ks.HashMap.hashFunc = hashFunc
	return ksb
}

func (ksb *KSBuilder) WithInitialHashMapUnderlyingArraySize(size int) *KSBuilder {
	ksb.ks.HashMap.bucket = make([][]*hashMapNode, size)
	return ksb
}

func (ksb *KSBuilder) WithListCapacity(capacity int) *KSBuilder {
	ksb.ks.List.capacity = capacity
	return ksb
}

func (ksb *KSBuilder) Build() *KS {
	return ksb.ks
}

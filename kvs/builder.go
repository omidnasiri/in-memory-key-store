package kvs

type KVSBuilder struct {
	kvs *KVS
}

func NewKVSBuilder() *KVSBuilder {
	return &KVSBuilder{
		kvs: newKVS(),
	}
}

func (b *KVSBuilder) WithHashFunction(hashFunc HashFunction) *KVSBuilder {
	b.kvs.hashMap.hashFunc = hashFunc
	return b
}

func (b *KVSBuilder) WithInitialHashMapUnderlyingArraySize(size int) *KVSBuilder {
	b.kvs.hashMap.bucket = make([][]*hashMapNode, size)
	return b
}

func (b *KVSBuilder) WithListCapacity(capacity int) *KVSBuilder {
	b.kvs.list.capacity = capacity
	return b
}

func (b *KVSBuilder) Build() *KVS {
	return b.kvs
}

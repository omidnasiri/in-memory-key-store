package kvs

const (
	fnvOffsetBasis uint64 = 14695981039346656037
	fnvPrime       uint64 = 1099511628211
)

// Universal hashing is the best implementation.
type HashFunction func(Key string, limit int) int

func SimpleHashFunc(key string, limit int) int {
	hash := 0
	for _, char := range key {
		hash += int(char)
	}
	return hash % limit
}

func FowlerNollVoHashFunction(key string, limit int) int {
	hash := fnvOffsetBasis
	byteArray := []byte(key)

	for i := range byteArray {
		hash ^= uint64(byteArray[i])
		hash *= fnvPrime
	}

	return int(hash % uint64(limit))
}

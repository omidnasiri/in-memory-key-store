package main

import "github.com/omidnasiri/in-memory-key-value-store/kvs"

func main() {
	// sample builder usage
	_ = kvs.NewKVSBuilder().
		WithHashFunction(kvs.FowlerNollVoHashFunction).
		WithInitialHashMapUnderlyingArraySize(5).
		WithListCapacity(5).
		Build()

	// pass kvs to application layer
}

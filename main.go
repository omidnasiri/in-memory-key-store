package main

import "github.com/omidnasiri/in-memory-key-store/ks"

func main() {
	// sample builder usage
	_ = ks.NewKSBuilder().
		WithHashFunction(ks.FowlerNollVoHashFunction).
		WithInitialHashMapUnderlyingArraySize(5).
		WithListCapacity(5).
		Build()
}

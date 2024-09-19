package main

import (
	"testing"
	"time"
)

func TestNewHashMap(t *testing.T) {
	initialSize := 5
	hm := NewHashMap(FowlerNollVoHashFunction, initialSize)

	if len(hm.bucket) != initialSize {
		t.Errorf("Expected bucket size %d, got %d", initialSize, len(hm.bucket))
	}

	if hm.hashFunc == nil {
		t.Errorf("Expected hash function to be assigned correctly")
	}

	if hm.population != 0 {
		t.Errorf("Expected population to be 0, got %d", hm.population)
	}
}

func TestSet(t *testing.T) {
	hm := NewHashMap(FowlerNollVoHashFunction, 5)
	key := "testKey"
	value := "testValue"
	ttl := 1 * time.Hour

	hm.Set(key, value, ttl)

	index := hm.hashFunc(key, len(hm.bucket))
	found := false
	for _, node := range hm.bucket[index] {
		if node.key == key {
			found = true
			if node.value != value {
				t.Errorf("Expected value %s, got %s", value, node.value)
			}
			if node.expiry.Before(time.Now().Add(ttl - time.Minute)) {
				t.Errorf("Expected expiry to be at least %v, got %v", ttl, node.expiry.Sub(time.Now()))
			}
		}
	}
	if !found {
		t.Errorf("Expected to find key %s in the hashmap", key)
	}
}

func TestGet(t *testing.T) {
	hm := NewHashMap(FowlerNollVoHashFunction, 5)
	key := "testKey"
	value := "testValue"
	ttl := 1 * time.Hour

	hm.Set(key, value, ttl)

	gotValue, found := hm.Get(key)
	if !found {
		t.Errorf("Expected to find key %s", key)
	}
	if gotValue != value {
		t.Errorf("Expected value %s, got %s", value, gotValue)
	}
}

func TestGetExpiredKey(t *testing.T) {
	hm := NewHashMap(FowlerNollVoHashFunction, 5)
	key := "testKey"
	value := "testValue"
	ttl := -1 * time.Hour // Set TTL in the past to expire immediately

	hm.Set(key, value, ttl)

	_, found := hm.Get(key)
	if found {
		t.Errorf("Expected not to find key %s as it should be expired", key)
	}
}

func TestSetOverwrite(t *testing.T) {
	hm := NewHashMap(FowlerNollVoHashFunction, 5)
	key := "testKey"
	value1 := "testValue1"
	value2 := "testValue2"
	ttl := 1 * time.Hour

	hm.Set(key, value1, ttl)
	hm.Set(key, value2, ttl)

	index := hm.hashFunc(key, len(hm.bucket))
	count := 0
	for _, node := range hm.bucket[index] {
		if node.key == key {
			count++
			if node.value != value2 {
				t.Errorf("Expected value %s, got %s", value2, node.value)
			}
		}
	}
	if count == 2 {
		t.Errorf("Expected to find only one key %s in the hashmap", key)
	}
}

func TestSetWithCollision(t *testing.T) {
	hm := NewHashMap(FowlerNollVoHashFunction, 1) // Force collision by using a single bucket
	key1 := "testKey1"
	value1 := "testValue1"
	key2 := "testKey2"
	value2 := "testValue2"
	ttl := 1 * time.Hour

	hm.Set(key1, value1, ttl)
	hm.Set(key2, value2, ttl)

	if len(hm.bucket[0]) != 2 {
		t.Errorf("Expected bucket to contain 2 items, got %d", len(hm.bucket[0]))
	}

	foundKey1 := false
	foundKey2 := false
	for _, node := range hm.bucket[0] {
		if node.key == key1 {
			foundKey1 = true
			if node.value != value1 {
				t.Errorf("Expected value %s for key %s, got %s", value1, key1, node.value)
			}
		}
		if node.key == key2 {
			foundKey2 = true
			if node.value != value2 {
				t.Errorf("Expected value %s for key %s, got %s", value2, key2, node.value)
			}
		}
	}
	if !foundKey1 {
		t.Errorf("Expected to find key %s in the hashmap", key1)
	}
	if !foundKey2 {
		t.Errorf("Expected to find key %s in the hashmap", key2)
	}
}

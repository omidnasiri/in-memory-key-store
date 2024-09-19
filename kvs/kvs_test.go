package kvs

import (
	"testing"
	"time"
)

func TestNewKVS(t *testing.T) {
	kvs := newKVS()

	if kvs == nil {
		t.Fatal("Expected NewKVS to return a non-nil KVS instance")
	}

	if kvs.hashMap == nil {
		t.Fatal("Expected HashMap to be initialized")
	}

	if kvs.list == nil {
		t.Fatal("Expected List to be initialized")
	}
}

func TestKVS_Set(t *testing.T) {
	kvs := newKVS()

	// Test setting a new key-value pair
	kvs.Set("key1", "value1", 10)
	value, ok := kvs.Get("key1")
	if !ok || value != "value1" {
		t.Fatalf("Expected key1 to have value 'value1', got '%s'", value)
	}

	// Test updating an existing key-value pair
	kvs.Set("key1", "value2", 10)
	value, ok = kvs.Get("key1")
	if !ok || value != "value2" {
		t.Fatalf("Expected key1 to have value 'value2', got '%s'", value)
	}

	// Test setting another key-value pair
	kvs.Set("key2", "value3", 10)
	value, ok = kvs.Get("key2")
	if !ok || value != "value3" {
		t.Fatalf("Expected key2 to have value 'value3', got '%s'", value)
	}

	// Test TTL expiration (assuming the TTL mechanism is implemented in HashMap)
	time.Sleep(10 * time.Second)
	_, ok = kvs.Get("key1")
	if ok {
		t.Fatal("Expected key1 to be expired and not found")
	}
}

func TestKVS_Get(t *testing.T) {
	kvs := newKVS()

	// Test getting a value for a non-existent key
	value, ok := kvs.Get("nonexistent")
	if ok {
		t.Fatalf("Expected 'nonexistent' key to not be found, but got value '%s'", value)
	}

	// Test getting a value for an existing key
	kvs.Set("key1", "value1", 10)
	value, ok = kvs.Get("key1")
	if !ok || value != "value1" {
		t.Fatalf("Expected key1 to have value 'value1', got '%s'", value)
	}

	// Test getting a value for a key that has expired (assuming TTL mechanism is implemented)
	kvs.Set("key2", "value2", 1) // 1 second TTL
	time.Sleep(2 * time.Second)  // Wait for the key to expire
	_, ok = kvs.Get("key2")
	if ok {
		t.Fatal("Expected key2 to be expired and not found")
	}
}

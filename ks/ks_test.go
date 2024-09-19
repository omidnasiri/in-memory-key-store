package ks

import (
	"testing"
	"time"
)

func TestNewKS(t *testing.T) {
	ks := NewKS()

	if ks == nil {
		t.Fatal("Expected NewKS to return a non-nil KS instance")
	}

	if ks.HashMap == nil {
		t.Fatal("Expected HashMap to be initialized")
	}

	if ks.List == nil {
		t.Fatal("Expected List to be initialized")
	}
}

func TestKS_Set(t *testing.T) {
	ks := NewKS()

	// Test setting a new key-value pair
	ks.Set("key1", "value1", 10)
	value, ok := ks.Get("key1")
	if !ok || value != "value1" {
		t.Fatalf("Expected key1 to have value 'value1', got '%s'", value)
	}

	// Test updating an existing key-value pair
	ks.Set("key1", "value2", 10)
	value, ok = ks.Get("key1")
	if !ok || value != "value2" {
		t.Fatalf("Expected key1 to have value 'value2', got '%s'", value)
	}

	// Test setting another key-value pair
	ks.Set("key2", "value3", 10)
	value, ok = ks.Get("key2")
	if !ok || value != "value3" {
		t.Fatalf("Expected key2 to have value 'value3', got '%s'", value)
	}

	// Test TTL expiration (assuming the TTL mechanism is implemented in HashMap)
	time.Sleep(10 * time.Second)
	_, ok = ks.Get("key1")
	if ok {
		t.Fatal("Expected key1 to be expired and not found")
	}
}

func TestKS_Get(t *testing.T) {
	ks := NewKS()

	// Test getting a value for a non-existent key
	value, ok := ks.Get("nonexistent")
	if ok {
		t.Fatalf("Expected 'nonexistent' key to not be found, but got value '%s'", value)
	}

	// Test getting a value for an existing key
	ks.Set("key1", "value1", 10)
	value, ok = ks.Get("key1")
	if !ok || value != "value1" {
		t.Fatalf("Expected key1 to have value 'value1', got '%s'", value)
	}

	// Test getting a value for a key that has expired (assuming TTL mechanism is implemented)
	ks.Set("key2", "value2", 1) // 1 second TTL
	time.Sleep(2 * time.Second) // Wait for the key to expire
	_, ok = ks.Get("key2")
	if ok {
		t.Fatal("Expected key2 to be expired and not found")
	}
}

package ks

import "testing"

func TestInsertHead(t *testing.T) {
	tests := []struct {
		name           string
		capacity       int
		insertions     []string
		expectedHead   string
		expectedTail   string
		expectedLength int
	}{
		{
			name:           "Insert into empty list",
			capacity:       3,
			insertions:     []string{"a"},
			expectedHead:   "a",
			expectedTail:   "a",
			expectedLength: 1,
		},
		{
			name:           "Insert into non-full list",
			capacity:       3,
			insertions:     []string{"a", "b"},
			expectedHead:   "b",
			expectedTail:   "a",
			expectedLength: 2,
		},
		{
			name:           "Insert into full list",
			capacity:       2,
			insertions:     []string{"a", "b", "c"},
			expectedHead:   "c",
			expectedTail:   "b",
			expectedLength: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := NewList(tt.capacity)
			var removedData string
			for _, data := range tt.insertions {
				removedData = list.InsertHead(data)
			}

			if list.head.data != tt.expectedHead {
				t.Errorf("expected head %v, got %v", tt.expectedHead, list.head.data)
			}

			if list.tail.data != tt.expectedTail {
				t.Errorf("expected tail %v, got %v", tt.expectedTail, list.tail.data)
			}

			if list.population != tt.expectedLength {
				t.Errorf("expected length %v, got %v", tt.expectedLength, list.population)
			}

			if tt.capacity == len(tt.insertions) && removedData != tt.insertions[0] {
				t.Errorf("expected removed data %v, got %v", tt.insertions[0], removedData)
			}
		})
	}
}
func TestNewList(t *testing.T) {
	tests := []struct {
		name     string
		capacity int
	}{
		{
			name:     "Create list with capacity 1",
			capacity: 1,
		},
		{
			name:     "Create list with capacity 5",
			capacity: 5,
		},
		{
			name:     "Create list with capacity 10",
			capacity: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := NewList(tt.capacity)

			if list.capacity != tt.capacity {
				t.Errorf("expected capacity %v, got %v", tt.capacity, list.capacity)
			}

			if list.head != nil {
				t.Errorf("expected head to be nil, got %v", list.head)
			}

			if list.tail != nil {
				t.Errorf("expected tail to be nil, got %v", list.tail)
			}

			if list.population != 0 {
				t.Errorf("expected population to be 0, got %v", list.population)
			}
		})
	}
}

func TestRemoveTail(t *testing.T) {
	tests := []struct {
		name            string
		capacity        int
		insertions      []string
		expectedTail    string
		expectedLength  int
		expectedRemoved string
	}{
		{
			name:            "Remove tail from empty list",
			capacity:        3,
			insertions:      []string{},
			expectedTail:    "",
			expectedLength:  0,
			expectedRemoved: "",
		},
		{
			name:            "Remove tail from single element list",
			capacity:        3,
			insertions:      []string{"a"},
			expectedTail:    "",
			expectedLength:  0,
			expectedRemoved: "a",
		},
		{
			name:            "Remove tail from multiple element list",
			capacity:        3,
			insertions:      []string{"a", "b", "c"},
			expectedTail:    "b",
			expectedLength:  2,
			expectedRemoved: "a",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := NewList(tt.capacity)
			for _, data := range tt.insertions {
				list.InsertHead(data)
			}

			removedData := list.removeTail()

			if list.tail != nil && list.tail.data != tt.expectedTail {
				t.Errorf("expected tail %v, got %v", tt.expectedTail, list.tail.data)
			}

			if list.population != tt.expectedLength {
				t.Errorf("expected length %v, got %v", tt.expectedLength, list.population)
			}

			if removedData != tt.expectedRemoved {
				t.Errorf("expected removed data %v, got %v", tt.expectedRemoved, removedData)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		name           string
		capacity       int
		insertions     []string
		deleteData     string
		expectedHead   string
		expectedTail   string
		expectedLength int
	}{
		{
			name:           "Delete from empty list",
			capacity:       3,
			insertions:     []string{},
			deleteData:     "a",
			expectedHead:   "",
			expectedTail:   "",
			expectedLength: 0,
		},
		{
			name:           "Delete head from single element list",
			capacity:       3,
			insertions:     []string{"a"},
			deleteData:     "a",
			expectedHead:   "",
			expectedTail:   "",
			expectedLength: 0,
		},
		{
			name:           "Delete head from multiple element list",
			capacity:       3,
			insertions:     []string{"a", "b", "c"},
			deleteData:     "c",
			expectedHead:   "b",
			expectedTail:   "a",
			expectedLength: 2,
		},
		{
			name:           "Delete tail from multiple element list",
			capacity:       3,
			insertions:     []string{"a", "b", "c"},
			deleteData:     "a",
			expectedHead:   "c",
			expectedTail:   "b",
			expectedLength: 2,
		},
		{
			name:           "Delete middle element from multiple element list",
			capacity:       3,
			insertions:     []string{"a", "b", "c"},
			deleteData:     "b",
			expectedHead:   "c",
			expectedTail:   "a",
			expectedLength: 2,
		},
		{
			name:           "Delete non-existent element",
			capacity:       3,
			insertions:     []string{"a", "b", "c"},
			deleteData:     "d",
			expectedHead:   "c",
			expectedTail:   "a",
			expectedLength: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := NewList(tt.capacity)
			for _, data := range tt.insertions {
				list.InsertHead(data)
			}

			list.Delete(tt.deleteData)

			if list.head != nil && list.head.data != tt.expectedHead {
				t.Errorf("expected head %v, got %v", tt.expectedHead, list.head.data)
			}

			if list.tail != nil && list.tail.data != tt.expectedTail {
				t.Errorf("expected tail %v, got %v", tt.expectedTail, list.tail.data)
			}

			if list.population != tt.expectedLength {
				t.Errorf("expected length %v, got %v", tt.expectedLength, list.population)
			}
		})
	}
}

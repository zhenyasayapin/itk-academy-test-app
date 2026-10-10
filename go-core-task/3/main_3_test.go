package main

import (
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		data map[string]int
	}{
		{nil},
		{make(map[string]int)},
	}

	var m *StringIntMap

	for _, tt := range tests {
		m = New(tt.data)

		if m == nil {
			t.Fatal("New() returned nil")
		}
		if m.data == nil {
			t.Error("New() returned map with nil data")
		}
	}
}

func TestAdd(t *testing.T) {
	s := New(make(map[string]int))

	s.Add("test", 1)
	value, exist := s.Get("test")

	if !exist {
		t.Error("Add() did not store the value")
	}
	if value != 1 {
		t.Errorf("Add() stored value = %d, want 1", value)
	}
}

func TestGet(t *testing.T) {
	tests := []struct {
		data  map[string]int
		key   string
		value int
		exist bool
	}{
		{map[string]int{"test": 123}, "test", 123, true},
		{map[string]int{}, "test", 0, false},
	}

	for _, tt := range tests {
		m := New(tt.data)

		value, exist := m.Get(tt.key)
		if exist != tt.exist {
			t.Errorf("Get(%q) exist = %v, want %v", tt.key, exist, tt.exist)
		}

		if value != tt.value {
			t.Errorf("Get(%q) = %d, want %d", tt.key, value, tt.value)
		}
	}

}

func TestRemove(t *testing.T) {
	tests := []struct {
		data map[string]int
		key  string
	}{
		{map[string]int{}, "test"},
		{map[string]int{"test": 2}, "test"},
	}

	for _, tt := range tests {
		s := New(tt.data)

		s.Remove(tt.key)
		actualValue, exist := s.data[tt.key]
		if exist == true {
			t.Error("Remove() key should not exist after removal")
		}

		if actualValue != 0 {
			t.Errorf("Remove() value = %d, want 0 (zero value)", actualValue)
		}
	}
}

func TestCopy(t *testing.T) {
	tests := []struct {
		data  map[string]int
		key   string
		value int
	}{
		{map[string]int{"test": 123}, "test", 456},
	}

	for _, tt := range tests {
		s := New(tt.data)

		c := s.Copy()

		c[tt.key] = tt.value

		originalValue, exist := s.Get(tt.key)

		if exist == false {
			t.Error("Copy() original map should still contain the key")
		}

		if originalValue == tt.value {
			t.Errorf("Copy() original value was modified to %d, want %d (independent copy)", originalValue, tt.value)
		}
	}
}

func TestExist(t *testing.T) {
	tests := []struct {
		data  map[string]int
		key   string
		value int
		exist bool
	}{
		{map[string]int{"test": 123}, "test", 123, true},
		{map[string]int{}, "test", 0, false},
	}

	for _, tt := range tests {
		m := New(tt.data)

		exist := m.Exists(tt.key)
		if exist != tt.exist {
			t.Errorf("Exists(%q) = %v, want %v", tt.key, exist, tt.exist)
		}
	}

}

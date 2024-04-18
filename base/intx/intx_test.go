package intx

import (
	"testing"
)

// TestInArray ...
func TestInArray(t *testing.T) {
	a := []int{1, 2, 3, 4}
	if !InArray(3, a) {
		t.Error()
	}
	if InArray(0, a) {
		t.Error()
	}
	if InArray(5, a) {
		t.Error()
	}
}

// TestInUnique ...
func TestUnique(t *testing.T) {
	a := []int{1, 2, 3, 4, 0, 0, 1, 3, 7, 8}
	if b := Unique(a); len(b) != 7 {
		t.Error()
	}
}

// TestFilter ...
func TestFilter(t *testing.T) {
	a := []int{1, 2, 3, 4, 0, 0}
	if b := Filter(a); len(b) != 4 {
		t.Error()
	}
}

// TestIntersect ...
func TestIntersect(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := []int{3, 4, 5, 6}
	if c := Intersect(a, b); len(c) != 2 {
		t.Error()
	}
}

// TestUnion ...
func TestUnion(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := []int{3, 4, 5, 6}
	if c := Union(a, b); len(c) != 6 {
		t.Error()
	}
}

// TestDiff ...
func TestDiff(t *testing.T) {
	a := []int{0, 1, 2, 3, 4}
	b := []int{3, 4, 5, 6}
	if c := Diff(a, b); len(c) != 3 {
		t.Error()
	}
}

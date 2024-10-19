package slices

import (
	"reflect"
	"testing"
)

func TestFilterIntegers(t *testing.T) {
	numbers := []int{2, 3, 4, 5, 6}
	f := func(i int) bool { return i > 2 }
	filtered := Filter(numbers, f)

	expected := []int{3, 4, 5, 6}
	if !reflect.DeepEqual(filtered, expected) {
		t.Errorf("Expected %v, got %v", expected, filtered)
	}
}

func TestFilterStrLen(t *testing.T) {
	names := []string{"Priscille", "Geoffrey", "Nala", "Kira"}
	f := func(s string) bool { return len(s) < 5 }
	filtered := Filter(names, f)

	expected := []string{"Nala", "Kira"}
	if !reflect.DeepEqual(filtered, expected) {
		t.Errorf("Expected %v, got %v", expected, filtered)
	}
}

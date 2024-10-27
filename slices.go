package slices

import (
	"fmt"
	"reflect"
)

// Map applies f to every element from s, and returns a new slice with the results.
// The initial slice s will not be updated.
// If any errors happen during the call of f(s[i]), Map will stop and return
// the error.
func Map[U, V any](s []U, f func(U) (V, error)) (mapped []V, err error) {
	for _, e := range s {
		newValue, err := f(e)
		if err != nil {
			return nil, fmt.Errorf("could not apply func to %v: %v", e, err)
		}

		mapped = append(mapped, newValue)
	}

	return
}

// Filter returns a new slice with the elements from s that satisfy the
// condition f(s[i]).
// The initial slice s will not be updated.
func Filter[U any](s []U, f func(U) bool) (mapped []U) {
	for _, e := range s {
		if f(e) {
			mapped = append(mapped, e)
		}
	}

	return
}

// CheckDuplicates returns true if any duplicates where found in s
func CheckDuplicates[U any](s []U) (bool) {
  for i1, e1 := range s {
    for i2, e2 := range s {
      if reflect.DeepEqual(e1, e2) && i1 != i2 {
        return true
      }
    }
  }

  return false
}

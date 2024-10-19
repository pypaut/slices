package slices

import "fmt"

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
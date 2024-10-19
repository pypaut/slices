package slices

func Filter[U any](s []U, f func(U) bool) (mapped []U) {
	for _, e := range s {
		if f(e) {
			mapped = append(mapped, e)
		}
	}

	return
}

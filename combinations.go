package itertools

import (
	"fmt"
	"iter"
)

// Generic version
func Combinations[T any](items []T, r int) iter.Seq[[]T] {

	return func(yield func([]T) bool) {
		n := len(items)
		if r > n || r < 0 {
			panic(fmt.Sprintf("invalid r value: permutating set of %v for subsets size %v", n, r))
		}

		indices := make([]int, r)
		for i := range r {
			indices[i] = i
		}

		if !yield(items[:r]) {
			return
		}

		for {
			i := r - 1
			for i >= 0 {
				if indices[i] != i+n-r {
					break
				}
				i--
			}
			if i == -1 {
				return
			}
			indices[i] += 1
			for j := i + 1; j < r; j++ {
				indices[j] = indices[j-1] + 1
			}
			temp := make([]T, r)
			for i := range r {
				temp[i] = items[indices[i]]
			}
			if !yield(temp) {
				return
			}

		}
	}
}

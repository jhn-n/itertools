package itertools

import (
	"fmt"
	"iter"
)

// https://docs.python.org/3/library/itertools.html#itertools.permutations
// Non-recursive method to generate all permutations of an integer slice.
// no speed advantage observed, but allows r < n subsets
func Permutations[T comparable](items []T, r int) iter.Seq[[]T] {

	return func(yield func([]T) bool) {

		n := len(items)
		if r > n || r < 0 {
			panic(fmt.Sprintf("invalid r value: permutating set of %v for subsets size %v", n, r))
		}

		indices := make([]int, n)
		for i := range indices {
			indices[i] = i
		}

		cycles := make([]int, r)
		for i := range r {
			cycles[i] = n - i
		}

		if !yield(items[:r]) {
			return
		}

		for n > 0 {
			reachedEnd := true
			for i := r - 1; i >= 0; i-- {
				cycles[i] -= 1
				if cycles[i] == 0 {
					temp := indices[i]
					for j := i; j < n-1; j++ {
						indices[j] = indices[j+1]
					}
					indices[n-1] = temp
					cycles[i] = n - i
				} else {
					j := cycles[i]
					indices[i], indices[n-j] = indices[n-j], indices[i]
					temp := make([]T, r)
					for i := range r {
						temp[i] = items[indices[i]]
					}
					if !yield(temp) {
						return
					}
					reachedEnd = false
					break
				}
			}
			if reachedEnd {
				return
			}
		}
	}
}

// Recursive function to generate all permutations of an integer slice.
func NPermutations[T comparable](items []T) iter.Seq[[]T] {

	return func(yield func([]T) bool) {
		permuteBacktrack(items, 0, yield)
	}
}

func permuteBacktrack[T comparable](items []T, start int, yield func([]T) bool) bool {
	if start == len(items) {
		// Create a copy of the current permutation to avoid mutation issues
		temp := make([]T, len(items))
		copy(temp, items)
		if !yield(temp) {
			return false
		}
		return true
	}

	no_early_exit := true
	for i := start; no_early_exit && i < len(items); i++ {
		// Swap the current element with the element at the start position
		items[start], items[i] = items[i], items[start]

		// Recursively find permutations for the rest of the slice
		no_early_exit = permuteBacktrack(items, start+1, yield)

		// Backtrack: swap back to restore the original order for the next iteration
		items[start], items[i] = items[i], items[start]
	}
	return no_early_exit
}

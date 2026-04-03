package itertools

import (
	"fmt"
	"iter"
)

// PermuteInt generates all permutations of an integer slice.
func PermuteInt(nums []int) iter.Seq[[]int] {

	return func(yield func([]int) bool) {
		permuteIntBacktrack(nums, 0, yield)
	}
}

func permuteIntBacktrack(nums []int, start int, yield func([]int) bool) bool {
	if start == len(nums) {
		// Create a copy of the current permutation to avoid mutation issues
		temp := make([]int, len(nums))
		copy(temp, nums)
		if !yield(temp) {
			return false
		}
		return true
	}

	normal_yield := true
	for i := start; i < len(nums); i++ {
		// Swap the current element with the element at the start position
		nums[start], nums[i] = nums[i], nums[start]

		// Recursively find permutations for the rest of the slice
		normal_yield = permuteBacktrack(nums, start+1, yield)

		// Backtrack: swap back to restore the original order for the next iteration
		nums[start], nums[i] = nums[i], nums[start]

		if !normal_yield {
			break
		}
	}
	return normal_yield
}

// Generic version
func Permute[T any](items []T) iter.Seq[[]T] {

	return func(yield func([]T) bool) {
		permuteBacktrack(items, 0, yield)
	}
}

func permuteBacktrack[T any](items []T, start int, yield func([]T) bool) bool {
	if start == len(items) {
		// Create a copy of the current permutation to avoid mutation issues
		temp := make([]T, len(items))
		copy(temp, items)
		if !yield(temp) {
			return false
		}
		return true
	}

	normal_yield := true
	for i := start; i < len(items); i++ {
		// Swap the current element with the element at the start position
		items[start], items[i] = items[i], items[start]

		// Recursively find permutations for the rest of the slice
		normal_yield = permuteBacktrack(items, start+1, yield)

		// Backtrack: swap back to restore the original order for the next iteration
		items[start], items[i] = items[i], items[start]

		if !normal_yield {
			break
		}
	}
	return normal_yield
}

// https://docs.python.org/3/library/itertools.html#itertools.permutations
// Alternative method to generate all permutations of an integer slice.
// no speed advantage observed, but allows r < n values

func PermuteInt2(nums []int, r int) iter.Seq[[]int] {

	return func(yield func([]int) bool) {

		n := len(nums)
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

		if !yield(nums[:r]) {
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
					temp := make([]int, n)
					for i := range r {
						temp[i] = nums[indices[i]]
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

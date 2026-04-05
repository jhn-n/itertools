package itertools

import (
	"iter"
	"slices"
	"testing"
)

var expectedNPermutations = [][][]int{
	{{}},
	{{1}},
	{{1, 2}, {2, 1}},
	{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 2, 1}, {3, 1, 2}},
}

func TestPermutations(t *testing.T) {
	for i, exp := range expectedNPermutations {
		args := buildSlice(i)
		testMatches(t, Permutations(args, len(args)), exp)
	}

	args := []string{"b", "a"}
	exp := [][]string{{"a", "b"}, {"b", "a"}}
	testMatches(t, Permutations(args, len(args)), exp)

	args = []string{"a", "b", "c"}

	r := 0
	exp = [][]string{{}}
	testMatches(t, Permutations(args, r), exp)

	r = 1
	exp = [][]string{{"a"}, {"b"}, {"c"}}
	testMatches(t, Permutations(args, r), exp)

	r = 2
	exp = [][]string{{"a", "b"}, {"b", "a"}, {"a", "c"},
		{"c", "a"}, {"b", "c"}, {"c", "b"}}
	testMatches(t, Permutations(args, r), exp)
}

func TestNPermutations(t *testing.T) {
	for i, exp := range expectedNPermutations {
		args := buildSlice(i)
		testMatches(t, NPermutations(args), exp)
	}

	args := []string{"b", "a"}
	exp := [][]string{{"a", "b"}, {"b", "a"}}
	testMatches(t, NPermutations(args), exp)
}

func TestPermutationsLen(t *testing.T) {
	for n := range 7 {
		arg := buildSlice(n)
		for r := 0; r <= n; r++ {
			got := 0
			for range Permutations(arg, r) {
				got += 1

			}
			exp := factorial(n) / factorial(n-r)
			if got != exp {
				t.Errorf("len(Permutations(%v)) = %v; want %v", arg, got, exp)
			}
		}
	}
}

func TestNPermutationsLen(t *testing.T) {
	for i := range 7 {
		arg := buildSlice(i)

		got := 0
		for range NPermutations(arg) {
			got += 1
		}
		exp := factorial(i)
		if got != exp {
			t.Errorf("len(NPermutations(%v)) = %v; want %v", arg, got, exp)
		}
	}
}

func BenchmarkPermutations(b *testing.B) {
	arg := []int{1, 2, 3, 4, 5, 6}
	for b.Loop() {
		for range Permutations(arg, 6) {
		}
	}
}

func BenchmarkNPermutations(b *testing.B) {
	arg := []int{1, 2, 3, 4, 5, 6}
	for b.Loop() {
		for range NPermutations(arg) {
		}
	}
}

func buildSlice(n int) []int {
	s := make([]int, n)
	for i := range n {
		s[i] = i + 1
	}
	return s
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func testMatches[T comparable](t *testing.T, seq iter.Seq[[]T], exp [][]T) {
	var got [][]T
	for p := range seq {
		got = append(got, p)
	}

	correct := true
	if len(exp) != len(got) {
		correct = false
	}

	for _, g := range got {
		found := false
		for _, e := range exp {
			if slices.Equal(g, e) {
				found = true
				break
			}
		}
		if !found {
			correct = false
			break
		}
	}

	for _, e := range exp {
		found := false
		for _, g := range got {
			if slices.Equal(g, e) {
				found = true
				break
			}
		}
		if !found {
			correct = false
			break
		}
	}

	if !correct {
		t.Errorf("got %v; want %v", got, exp)
	}
}

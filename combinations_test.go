package itertools

import (
	"fmt"
	"slices"
	"testing"
)

type testType = struct {
	items []int
	r     int
	exp   [][]int
}

var testSuite = []testType{
	{[]int{}, 0, [][]int{{}}},
	{[]int{1}, 1, [][]int{{1}}},
	{[]int{1, 2}, 1, [][]int{{1}, {2}}},
	{[]int{1, 2}, 2, [][]int{{1, 2}}},
	{[]int{1, 2, 3}, 0, [][]int{{}}},
	{[]int{1, 2, 3}, 1, [][]int{{1}, {2}, {3}}},
	{[]int{1, 2, 3}, 2, [][]int{{1, 2}, {1, 3}, {2, 3}}},
	{[]int{1, 2, 3}, 3, [][]int{{1, 2, 3}}},
	{[]int{1, 2, 3, 4}, 2, [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}},
}

func TestCombinations(t *testing.T) {
	for _, test := range testSuite {
		var exp []string
		for _, c := range test.exp {
			exp = append(exp, fmt.Sprint(c))
		}
		slices.Sort(exp)

		var got []string
		for c := range Combinations(test.items, test.r) {
			got = append(got, fmt.Sprint(c))
		}
		slices.Sort(got)

		if fmt.Sprint(exp) != fmt.Sprint(got) {
			t.Errorf("PermuteInt(%v,%v) = %v; want %v", test.items, test.r, got, exp)
			break
		}
	}
}

func TestCombinationLen(t *testing.T) {
	for i := 5; i < 11; i++ {
		arg := buildSlice(i)

		got := 0
		for range Combinations(arg, 3) {
			got += 1
		}
		exp := nChoosek(i, 3)
		if got != exp {
			t.Errorf("len(PermuteInt(%v)) = %v; want %v", arg, got, exp)
		}
	}
}

func BenchmarkCombination(b *testing.B) {
	arg := []int{1, 2, 3, 4, 5, 6}
	for b.Loop() {
		for range Combinations(arg, 3) {
		}
	}
}

func nChoosek(n, k int) int {
	return factorial(n) / factorial(k) / factorial(n-k)
}

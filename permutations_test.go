package itertools

import (
	"fmt"
	"slices"
	"testing"
)

var expectedPerms = [][][]int{
	{{}},
	{{1}},
	{{1, 2}, {2, 1}},
	{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
}

func TestPermuteInt(t *testing.T) {
	for i, perms := range expectedPerms {
		arg := buildSlice(i)

		var exp []string
		for _, p := range perms {
			exp = append(exp, fmt.Sprint(p))
		}
		slices.Sort(exp)

		var got []string
		for p := range PermuteInt(arg) {
			got = append(got, fmt.Sprint(p))
		}
		slices.Sort(got)

		if fmt.Sprint(exp) != fmt.Sprint(got) {
			t.Errorf("PermuteInt(%v) = %v; want %v", arg, got, exp)
			break
		}
	}
}

func TestPermute(t *testing.T) {
	for i, perms := range expectedPerms {
		arg := buildSlice(i)

		var exp []string
		for _, p := range perms {
			exp = append(exp, fmt.Sprint(p))
		}
		slices.Sort(exp)

		var got []string
		for p := range Permute(arg) {
			got = append(got, fmt.Sprint(p))
		}
		slices.Sort(got)

		if fmt.Sprint(exp) != fmt.Sprint(got) {
			t.Errorf("PermuteInt(%v) = %v; want %v", arg, got, exp)
			break
		}
	}

	arg2 := []string{"b", "a"}
	exp2 := [][]string{{"a", "b"}, {"b", "a"}}

	var got2 []string
	for p := range Permute(arg2) {
		got2 = append(got2, fmt.Sprint(p))
	}
	slices.Sort(got2)
	if fmt.Sprint(exp2) != fmt.Sprint(got2) {
		t.Errorf("PermuteInt(%v) = %v; want %v", arg2, got2, exp2)
	}

}

func TestPermuteInt2(t *testing.T) {
	for i, perms := range expectedPerms {
		arg := buildSlice(i)

		var exp []string
		for _, p := range perms {
			exp = append(exp, fmt.Sprint(p))
		}
		slices.Sort(exp)

		var got []string
		for p := range PermuteInt2(arg, len(arg)) {
			got = append(got, fmt.Sprint(p))
		}
		slices.Sort(got)

		if fmt.Sprint(exp) != fmt.Sprint(got) {
			t.Errorf("PermuteInt(%v) = %v; want %v", arg, got, exp)
			break
		}
	}
}

func TestPermuteIntLen(t *testing.T) {
	for i := range 7 {
		arg := buildSlice(i)

		got := 0
		for range PermuteInt(arg) {
			got += 1
		}
		exp := factorial(i)
		if got != exp {
			t.Errorf("len(PermuteInt(%v)) = %v; want %v", arg, got, exp)
		}
	}
}

func TestPermuteLen(t *testing.T) {
	for i := range 7 {
		arg := buildSlice(i)

		got := 0
		for range Permute(arg) {
			got += 1
		}
		exp := factorial(i)
		if got != exp {
			t.Errorf("len(PermuteInt(%v)) = %v; want %v", arg, got, exp)
		}
	}
}

func TestPermuteInt2Len(t *testing.T) {
	for i := range 7 {
		arg := buildSlice(i)

		got := 0
		for range PermuteInt2(arg, len(arg)) {
			got += 1
		}
		exp := factorial(i)
		if got != exp {
			t.Errorf("len(PermuteInt(%v)) = %v; want %v", arg, got, exp)
		}
	}
}

func BenchmarkPermuteInt(b *testing.B) {
	arg := []int{1, 2, 3, 4, 5, 6}
	for b.Loop() {
		for range PermuteInt(arg) {
		}
	}
}

func BenchmarkPermute(b *testing.B) {
	arg := []int{1, 2, 3, 4, 5, 6}
	for b.Loop() {
		for range Permute(arg) {
		}
	}
}

func BenchmarkPermuteInt2(b *testing.B) {
	arg := []int{1, 2, 3, 4, 5, 6}
	for b.Loop() {
		for range PermuteInt2(arg, 6) {
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

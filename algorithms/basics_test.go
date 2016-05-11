package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	Reverse(nums)
	assert.Equal(t, nums, []int{5, 4, 3, 2, 1})
}

// need to shuffle N times, bin results, and ensure
// uniform distribution via Pearson Chi-squared test
func TestShuffle(t *testing.T) {
	t.Parallel()

	n := 4
	iterations := 100000
	var bins = make(map[string]int)
	var nums = make([]int, 0, n)

	for i := 0; i <= n; i++ {
		nums = append(nums, i)
	}

	for i := 0; i <= iterations; i++ {
		tmp := make([]int, n)
		copy(tmp, nums)

		Shuffle(tmp)
		bins[fmt.Sprint(tmp)]++
	}

	t.Skip("Please implement Pearson Chi-squared test")
}

func TestBinarySearch(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}

	cases := []struct {
		Haystack []int
		Needle   int
		Expected int
	}{
		{nums, 1, 0},
		{nums, 3, 2},
		{nums, 5, 4},
		{nums, 0, -1},
		{nums, 6, -1},
		{[]int{}, 6, -1},
	}

	for _, tc := range cases {
		actual := BinarySearch(tc.Haystack, tc.Needle)
		actualRecursive := BinarySearchRecursive(tc.Haystack, tc.Needle)

		assert.Equal(t, tc.Expected, actual)
		assert.Equal(t, tc.Expected, actualRecursive)
	}
}

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		String   string
		Expected bool
	}{
		{"level", true},
		{"sagas", true},
		{"101", true},
		{"A", true},
		{"a", true},
		{"aa", true},
		{"", true},
		{"abc", false},
		{"sally", false},
	}

	for _, tc := range cases {
		actual := IsPalindrome(tc.String)
		actualRecursive := IsPalindromeRecursive(tc.String)

		assert.Equal(t, tc.Expected, actual)
		assert.Equal(t, tc.Expected, actualRecursive)
	}
}

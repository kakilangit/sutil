package sutil_test

import (
	"math"
	"sort"
	"testing"

	"github.com/kakilangit/sutil"
)

func TestSplit(t *testing.T) {
	t.Parallel()

	for name, test := range map[string]struct {
		input          []string
		limit          int
		expectedLength int
		expectedTotal  int
		err            error
	}{
		"given limit is zero then returns error": {
			input: []string{""},
			err:   sutil.ErrInvalidLimit,
		},
		"given limit negative then returns error": {
			input: []string{""},
			limit: -1,
			err:   sutil.ErrInvalidLimit,
		},
		"given limit is more than maximum integer then returns error": {
			input: []string{""},
			limit: math.MaxInt32 + 1,
			err:   sutil.ErrInvalidLimit,
		},
		"given nil_slice then returns error": {
			input: nil,
			limit: 1,
			err:   sutil.ErrInvalidStringSlice,
		},
		"given slice of 3 strings with limit 1 then expect length 3 and total 3": {
			input:          []string{"A", "B", "C"},
			limit:          1,
			expectedLength: 3,
			expectedTotal:  3,
		},
		"given slice of 3 strings with limit 2 then expect length 2 and total 3": {
			input:          []string{"A", "B", "C"},
			limit:          2,
			expectedLength: 2,
			expectedTotal:  3,
		},
		"given slice of 9 strings with limit 5 then expect length 2 and total 9": {
			input:          []string{"A", "B", "C", "D", "E", "F", "G", "H", "I"},
			limit:          5,
			expectedLength: 2,
			expectedTotal:  9,
		},
		"given slice of 10 strings with limit 3 then expect length 4 and total 10": {
			input:          []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"},
			limit:          3,
			expectedLength: 4,
			expectedTotal:  10,
		},
	} {
		t.Run(name, func(t *testing.T) {
			output, err := sutil.Split(test.input, test.limit)
			if test.err != err {
				t.Error("error must be equal")
			}

			if test.expectedLength != len(output) {
				t.Error("slice length must be equal")
			}

			total := 0
			for i := range output {
				total += len(output[i])
			}

			if test.expectedTotal != total {
				t.Error("total member of the slice must be equal")
			}
		})
	}
}

func TestTotalPage(t *testing.T) {
	t.Parallel()

	for name, test := range map[string]struct {
		limit         int
		slices        []int
		expectedTotal int
	}{
		"given slice of 2 integers and limit 1 expect total of 2": {
			limit:         1,
			slices:        make([]int, 2),
			expectedTotal: 2,
		},
		"given slice of 8 integers and limit 10 expect total of 1": {
			limit:         10,
			slices:        make([]int, 8),
			expectedTotal: 1,
		},
		"given slice of 8 integers and limit 3 expect total of 3": {
			limit:         3,
			slices:        make([]int, 8),
			expectedTotal: 3,
		},
		"given slice of 8 integers and limit 1 expect total of 8": {
			limit:         1,
			slices:        make([]int, 8),
			expectedTotal: 8,
		},
	} {
		t.Run(name, func(t *testing.T) {
			total := sutil.TotalPage(test.slices, test.limit)
			if test.expectedTotal != total {
				t.Error("total must be equal with the expected values")
			}
		})
	}
}

func TestUnique(t *testing.T) {
	t.Parallel()

	for name, test := range map[string]struct {
		slices   []int
		expected []int
	}{
		"given slice of 2 same integers expect total of 1": {
			slices:   []int{1, 1},
			expected: []int{1},
		},
		"given slice of 2 different integers expect total of 2": {
			slices:   []int{1, 2},
			expected: []int{1, 2},
		},
		"given slice of 5 mixed integers expect total of 4": {
			slices:   []int{1, 2, 2, 3, 10},
			expected: []int{1, 2, 3, 10},
		},
	} {
		t.Run(name, func(t *testing.T) {
			result := sutil.Unique(test.slices)

			sort.Ints(test.expected)
			sort.Ints(result)

			if !sutil.Equal(result, test.expected) {
				t.Errorf("result must be equal with the expected values")
			}
		})
	}

}

func BenchmarkSplit(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = sutil.Split([]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}, 3)
	}
}

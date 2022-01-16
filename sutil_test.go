package sutil_test

import (
	"fmt"
	"math"
	"sort"
	"testing"

	"github.com/kakilangit/sutil"
)

func tFataf(t *testing.T, msg string, expected, got interface{}) {
	t.Fatalf(`%s expected: %+v, got: %+v`, msg, expected, got)
}

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
			err:   sutil.ErrInvalidSlice,
		},
		"given empty slice then expect empty slice": {
			input: []string{},
			limit: 1,
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
				tFataf(t, "error must be equal", test.err, err)
			}

			if test.expectedLength != len(output) {
				tFataf(t, "slice length must be equal", test.expectedLength, len(output))
			}

			total := 0
			for i := range output {
				total += len(output[i])
			}

			if test.expectedTotal != total {
				tFataf(t, "total member of the slice must be equal", test.expectedTotal, total)
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
				tFataf(t, "total must be equal with the expected values", test.expectedTotal, total)
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
				tFataf(t, "result must be equal with the expected values", test.expected, result)
			}
		})
	}
}

func TestMap(t *testing.T) {
	t.Parallel()

	t.Run("given slice of integers and apply string mapper then expect slice of strings", func(t *testing.T) {
		given := []int{3, 2, 1}
		expected := []string{"0-3", "1-2", "2-1"}
		mapper := func(i, val int) string {
			return fmt.Sprintf(`%d-%d`, i, val)
		}

		result := sutil.Map(given, mapper)
		if !sutil.Equal(result, expected) {
			tFataf(t, "result must be equal with the expected values", expected, result)
		}
	})

	t.Run("given slice of float and apply integer mapper then expect slice of integer", func(t *testing.T) {
		given := []float64{3.2, 2.9, 1.1}
		expected := []int{3, 3, 1}
		mapper := func(_ int, val float64) int {
			return int(math.Round(val))
		}

		result := sutil.Map(given, mapper)
		if !sutil.Equal(result, expected) {
			tFataf(t, "result must be equal with the expected values", expected, result)
		}
	})
}

func TestReduce(t *testing.T) {
	t.Parallel()

	t.Run("given slice of integers and apply string reducer then expect string", func(t *testing.T) {
		given := []int{3, 2, 1}
		expected := "321"
		reducer := func(current string, val int) string {
			return fmt.Sprintf(`%s%d`, current, val)
		}

		result := sutil.Reduce(given, "", reducer)
		if result != expected {
			tFataf(t, "result must be equal with the expected values", expected, result)
		}
	})

	t.Run("given slice of float and apply integer reducer then expect integer", func(t *testing.T) {
		given := []float64{3.2, 2.9, 1.1}
		expected := 7
		reducer := func(current int, val float64) int {
			return current + int(math.Round(val))
		}

		result := sutil.Reduce(given, 0, reducer)
		if result != expected {
			tFataf(t, "result must be equal with the expected values", expected, result)
		}
	})
}

func TestFilter(t *testing.T) {
	t.Parallel()

	t.Run("given slice of integers and apply filter then expect filtered slice of integers", func(t *testing.T) {
		given := []int{3, 2, 1}
		expected := []int{2}
		filter := func(_, val int) bool {
			return val%2 == 0
		}

		result := sutil.Filter(given, filter)
		if !sutil.Equal(result, expected) {
			tFataf(t, "result must be equal with the expected values", expected, result)
		}
	})

	t.Run("given slice of strings and apply filter then expect filtered slice of strings", func(t *testing.T) {
		given := []string{"a", "b", "c"}
		expected := []string{"b", "c"}
		filter := func(_ int, val string) bool {
			return val != "a"
		}

		result := sutil.Filter(given, filter)
		if !sutil.Equal(result, expected) {
			tFataf(t, "result must be equal with the expected values", expected, result)
		}
	})
}

func BenchmarkSplit(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = sutil.Split([]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}, 3)
	}
}

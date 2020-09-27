package sutil_test

import (
	"math"
	"testing"

	"github.com/kakilangit/sutil"
)

func TestStringsSplit(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name           string
		input          []string
		limit          int
		expectedLength int
		expectedTotal  int
		err            error
	}{
		{
			name:  "zero_limit",
			input: []string{""},
			err:   sutil.ErrInvalidLimit,
		},
		{
			name:  "negative_limit",
			input: []string{""},
			limit: -1,
			err:   sutil.ErrInvalidLimit,
		},
		{
			name:  "max_limit",
			input: []string{""},
			limit: math.MaxInt32 + 1,
			err:   sutil.ErrInvalidLimit,
		},
		{
			name:  "nil_slice",
			input: nil,
			limit: 1,
			err:   sutil.ErrInvalidStringSlice,
		},
		{
			name:           "ok_1",
			input:          []string{"A", "B", "C"},
			limit:          1,
			expectedLength: 3,
			expectedTotal:  3,
		},
		{
			name:           "ok_2",
			input:          []string{"A", "B", "C"},
			limit:          2,
			expectedLength: 2,
			expectedTotal:  3,
		},
		{
			name:           "ok_3",
			input:          []string{"A", "B", "C", "D", "E", "F", "G", "H", "I"},
			limit:          5,
			expectedLength: 2,
			expectedTotal:  9,
		},
		{
			name:           "ok_4",
			input:          []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"},
			limit:          3,
			expectedLength: 4,
			expectedTotal:  10,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output, err := sutil.SplitStrings(test.input, test.limit)
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

func TestIndex(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name                       string
		page, limit, total         int
		expectedStart, expectedEnd int
	}{
		{
			name:          "ok_1",
			page:          0,
			limit:         1,
			total:         2,
			expectedStart: 0,
			expectedEnd:   1,
		},
		{
			name:          "ok_2",
			page:          0,
			limit:         10,
			total:         8,
			expectedStart: 0,
			expectedEnd:   8,
		},
		{
			name:          "ok_3",
			page:          1,
			limit:         3,
			total:         8,
			expectedStart: 3,
			expectedEnd:   6,
		},
		{
			name:          "ok_4",
			page:          10,
			limit:         1,
			total:         8,
			expectedStart: 8,
			expectedEnd:   8,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			start, end := sutil.Index(test.page, test.limit, test.total)
			if test.expectedStart != start || test.expectedEnd != end {
				t.Error("start & end index must be equal with the expected values")
			}
		})
	}
}

func TestTotalPage(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		limit, total  int
		expectedTotal int
	}{
		{
			name:          "ok_1",
			limit:         1,
			total:         2,
			expectedTotal: 2,
		},
		{
			name:          "ok_2",
			limit:         10,
			total:         8,
			expectedTotal: 1,
		},
		{
			name:          "ok_3",
			limit:         3,
			total:         8,
			expectedTotal: 3,
		},
		{
			name:          "ok_4",
			limit:         1,
			total:         8,
			expectedTotal: 8,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			total := sutil.TotalPage(test.limit, test.total)
			if test.expectedTotal != total {
				t.Error("total must be equal with the expected values")
			}
		})
	}
}

func BenchmarkStringsSplit(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = sutil.SplitStrings([]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}, 3)
	}
}

func BenchmarkIndex(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = sutil.Index(1, 2, 3)
	}
}

func BenchmarkTotalPage(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = sutil.TotalPage(2, 3)
	}
}

package sutil_test

import (
	"math"
	"testing"

	"github.com/kakilangit/sutil"
	"github.com/stretchr/testify/assert"
)

func TestStringsSplit(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name          string
		input         []string
		limit         int
		expectLength  int
		expectedTotal int
		err           error
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
			name:          "ok_1",
			input:         []string{"A", "B", "C"},
			limit:         1,
			expectLength:  3,
			expectedTotal: 3,
		},
		{
			name:          "ok_2",
			input:         []string{"A", "B", "C"},
			limit:         2,
			expectLength:  2,
			expectedTotal: 3,
		},
		{
			name:          "ok_3",
			input:         []string{"A", "B", "C", "D", "E", "F", "G", "H", "I"},
			limit:         5,
			expectLength:  2,
			expectedTotal: 9,
		},
		{
			name:          "ok_4",
			input:         []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"},
			limit:         3,
			expectLength:  4,
			expectedTotal: 10,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output, err := sutil.StringsSplit(test.input, test.limit)
			assert.Equal(t, test.err, err)
			assert.Len(t, output, test.expectLength)

			total := 0
			for i := range output {
				total += len(output[i])
			}

			assert.Equal(t, test.expectedTotal, total)
		})
	}
}

func BenchmarkStringsSplit(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = sutil.StringsSplit([]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}, 3)
	}
}

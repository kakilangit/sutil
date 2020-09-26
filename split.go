package sutil

import (
	"math"
)

// StringsSplit splits slice of string into slice of string slice with maximum member of string slice is limit
// Can be used for pagination or break the parameters used in SQL IN statements
// Manipulating slice is faster than append via iteration
func StringsSplit(s []string, limit int) ([][]string, error) {
	if limit < 1 || limit > math.MaxInt32 {
		return nil, ErrInvalidLimit
	}

	total := len(s)
	if s == nil && total == 0 {
		return nil, ErrInvalidStringSlice
	}

	slices := make([][]string, 0)
	for page := 1; page <= totalPage(total, limit); page++ {
		start, end := index(page, limit, total)
		slices = append(slices, s[start:end])
	}

	return slices, nil
}

func index(page, limit, total int) (int, int) {
	start := (page - 1) * limit
	if start > total {
		start = total
	}

	end := start + limit
	if end > total {
		end = total
	}

	return start, end
}

func totalPage(total, limit int) int {
	var (
		page   = total / limit
		remain = total % limit
	)

	if remain > 0 {
		page += 1
	}

	return page
}
package sutil

import (
	"math"
)

// SplitStrings splits slice of string into slice of string slice with maximum member of string slice is limit.
// Can be used for pagination or break the parameters used in SQL IN statements.
// Manipulating slice is faster than append via iteration.
//
// Waiting for Go 2 generic and all Split{Type}s will be renamed to Split that accept []T
//
// Example:
//	input := []string{"7892141641500", "7892141600279", "7892141600422", "7892141640145", "7892141650236", "7892141650274", "7892141650311"}
//	limit := 2
//
//	pages, err := sutil.SplitStrings(input, limit)
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(pages)
//
// Will return:
//
// 	[[7892141641500 7892141600279] [7892141600422 7892141640145] [7892141650236 7892141650274] [7892141650311]]
func SplitStrings(s []string, limit int) ([][]string, error) {
	if limit < 1 || limit > math.MaxInt32 {
		return nil, ErrInvalidLimit
	}

	length := len(s)
	if s == nil && length == 0 {
		return nil, ErrInvalidStringSlice
	}

	total := TotalPage(limit, length)
	slices := make([][]string, total)
	for page := 0; page < total; page++ {
		start, end := Index(page, limit, length)
		slices[page] = s[start:end]
	}

	return slices, nil
}

// Index is taking page, limit, and slice length and return the correct start and end index of slice
func Index(page, limit, length int) (int, int) {
	start := page * limit
	if start > length {
		start = length
	}

	end := start + limit
	if end > length {
		end = length
	}

	return start, end
}

// TotalPage is taking limit and slice length and return the total page for any given slice,
// instead of count the slice it required calculated length, making it more reusable for any type of slices.
//
// Waiting Go 2 generic to change the signature to accept []T and perform length calculation.
func TotalPage(limit, length int) int {
	var (
		page   = length / limit
		remain = length % limit
	)

	if remain > 0 {
		page += 1
	}

	return page
}

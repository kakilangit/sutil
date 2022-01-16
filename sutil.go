package sutil

import "math"

// Split splits []T into []T with maximum member of[]T is limit.
//
// Can be used for pagination or break the parameters used in SQL IN statements.
// Manipulating slice is faster than append via iteration.
//
// Example:
//	input := []string{"7892141641500", "7892141600279", "7892141600422", "7892141640145", "7892141650236", "7892141650274", "7892141650311"}
//	limit := 2
//
//	pages, err := sutil.Split(input, limit)
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(pages)
//
// Will return:
//
// 	[[7892141641500 7892141600279] [7892141600422 7892141640145] [7892141650236 7892141650274] [7892141650311]]
func Split[T any](ss []T, limit int) ([][]T, error) {
	if ss == nil {
		return nil, ErrInvalidSlice
	}

	if limit < 1 || limit > math.MaxInt32 {
		return nil, ErrInvalidLimit
	}

	total := TotalPage(ss, limit)
	slices := make([][]T, total)
	for page := 0; page < total; page++ {
		slices[page] = Content(ss, page, limit)
	}

	return slices, nil
}

// TotalPage returns the total page for any given slice based on the limit provided.
func TotalPage[T any](ss []T, limit int) int {
	var (
		length = len(ss)
		page   = length / limit
		remain = length % limit
	)

	if remain > 0 {
		page += 1
	}

	return page
}

// Content returns the content of the []T based on page and limit.
func Content[T any](ss []T, page, limit int) []T {
	length := len(ss)
	start := page * limit
	if start > length {
		start = length
	}

	end := start + limit
	if end > length {
		end = length
	}

	return ss[start:end]
}

// Map turns a []T to a []U using a mapping function.
func Map[T, U any](ss []T, fn func(int, T) U) []U {
	out := make([]U, len(ss))
	for i, v := range ss {
		out[i] = fn(i, v)
	}

	return out
}

// Reduce reduces a []T to a single value using a reduction function.
func Reduce[T, U any](ss []T, initializer U, f func(U, T) U) U {
	out := initializer
	for _, v := range ss {
		out = f(out, v)
	}

	return out
}

// Filter filters values from a slice using a filter function.
func Filter[T any](ss []T, fn func(int, T) bool) []T {
	out := make([]T, 0)
	for i, v := range ss {
		if !fn(i, v) {
			continue
		}

		out = append(out, v)
	}

	return out
}

// Unique make []T unique.
func Unique[T comparable](ss []T) []T {
	check := make(map[T]struct{})
	for _, s := range ss {
		check[s] = struct{}{}
	}

	res := make([]T, 0, len(check))
	for s := range check {
		res = append(res, s)
	}

	return res
}

// Equal check slice equality.
func Equal[T comparable](ss, compared []T) bool {
	if len(ss) != len(compared) {
		return false
	}

	for i, v := range ss {
		if v != compared[i] {
			return false
		}
	}

	return true
}

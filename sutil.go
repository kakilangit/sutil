package sutil

import (
	"errors"
	"fmt"
	"math"
	"reflect"
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

type StructSlice []interface{}

// NoAllocStringMap will generate map[string]struct{} of string based on struct field name
// This kind of hashmap is useful for uniqueness
func (s StructSlice) NoAllocStringMap(name string) (map[string]struct{}, error) {
	if s.isEmpty() {
		return nil, nil
	}

	var res = make(map[string]struct{})

	for i := range s {
		if !s.isValid(i) {
			return nil, errors.New("invalid type of struct")
		}

		val, err := s.getFieldStringValue(i, name)
		if err != nil {
			return nil, err
		}
		res[val] = struct{}{}
	}

	return res, nil
}

// StringSlice will generate slice of string based on struct field name
func (s StructSlice) StringSlice(name string) ([]string, error) {
	if s.isEmpty() {
		return nil, nil
	}

	var list []string

	for i := range s {
		if !s.isValid(i) {
			return nil, errors.New("invalid type of struct")
		}

		val, err := s.getFieldStringValue(i, name)
		if err != nil {
			return nil, err
		}
		list = append(list, val)
	}

	return list, nil
}

// StringSliceUnique will generate unique slice of string based on struct field name
func (s StructSlice) StringSliceUnique(name string) ([]string, error) {
	if s.isEmpty() {
		return nil, nil
	}

	unique, err := s.NoAllocStringMap(name)
	if err != nil {
		return nil, err
	}

	var (
		list = make([]string, len(unique))
		i    = 0
	)

	for k := range unique {
		list[0] = k
		i++
	}

	return list, nil
}

func (s StructSlice) isValid(i int) bool {
	obj := s[i]
	types := []reflect.Kind{reflect.Struct, reflect.Ptr}
	for _, t := range types {
		if reflect.TypeOf(obj).Kind() == t {
			return true
		}
	}

	return false
}

func (s StructSlice) isEmpty() bool {
	return len(s) == 0
}

func (s StructSlice) getValue(i int) reflect.Value {
	var (
		val reflect.Value
		obj = s[i]
	)

	switch reflect.TypeOf(obj).Kind() {
	case reflect.Ptr:
		val = reflect.ValueOf(obj).Elem()
	default:
		val = reflect.ValueOf(obj)
	}

	return val
}

func (s StructSlice) getFieldStringValue(i int, name string) (string, error) {
	field := s.getValue(i).FieldByName(name)
	if !field.IsValid() {
		return "", fmt.Errorf("field %s not exists", name)
	}
	val, ok := field.Interface().(string)
	if !ok {
		return "", fmt.Errorf("field %s is not string", name)
	}
	return val, nil
}

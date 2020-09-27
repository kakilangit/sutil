package sutil

// Error type will conform errorer interface and use it as error constant
type Error string

// Error method to conform errorer interface
func (e Error) Error() string {
	return string(e)
}

const (
	ErrInvalidStringSlice = Error("invalid string slice") // Invalid string slice
	ErrInvalidLimit       = Error("invalid limit value")  // Invalid limit
)

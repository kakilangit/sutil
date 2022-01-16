package sutil

// Error type will conform errorer interface and use it as error constant.
type Error string

// Error method to conform errorer interface.
func (e Error) Error() string {
	return string(e)
}

const (
	// ErrInvalidSlice is invalid string slice
	ErrInvalidSlice = Error("invalid slice")
	// ErrInvalidLimit is invalid limit
	ErrInvalidLimit = Error("invalid limit value")
)

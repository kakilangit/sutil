package sutil

// Error type will conform errorer interface and use it as error constant.
type Error string

// Error method to conform errorer interface.
func (e Error) Error() string {
	return string(e)
}

const (
	// ErrInvalidStringSlice is invalid string slice
	ErrInvalidStringSlice = Error("invalid string slice")
	// ErrInvalidLimit is invalid limit
	ErrInvalidLimit = Error("invalid limit value")
)

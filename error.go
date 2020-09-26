package sutil

type Error string

func (e Error) Error() string {
	return string(e)
}

const (
	ErrInvalidStringSlice = Error("invalid string slice")
	ErrInvalidLimit       = Error("invalid limit value")
)

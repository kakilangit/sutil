package sutil_test

import (
	"testing"

	"github.com/kakilangit/sutil"
)

func TestError_Error(t *testing.T) {
	t.Parallel()

	const expectedError = "error"
	var err error = sutil.Error(expectedError)
	if err == nil {
		t.Error("error must not be nil")
	}

	if err.Error() != expectedError {
		t.Error("error value must be equal")
	}
}

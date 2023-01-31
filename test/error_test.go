package test

import (
	"errors"
	"testing"
)

func TestError(t *testing.T) {
	err := errors.New("123")
	if err.Error() != "123" {
		t.Fail()
	}
}

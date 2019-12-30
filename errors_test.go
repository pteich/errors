package errors

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	testError := New("my error")

	assert.Error(t, testError)
	assert.Equal(t, "my error", testError.Error())
}

func TestWrap(t *testing.T) {
	testError := Wrap(os.ErrNotExist, "my error")

	assert.Error(t, testError)
	assert.True(t, errors.Is(testError, os.ErrNotExist))
	assert.Equal(t, "my error: file does not exist", testError.Error())
}

func TestErrorf(t *testing.T) {
	format := "could not open file %s"
	testError := Errorf(format, "foo")

	assert.Error(t, testError)
	assert.Equal(t, "could not open file foo", testError.Error())
}

func TestWrapf(t *testing.T) {
	format := "could not open file %s"
	testError := Wrapf(os.ErrNotExist, format, "foo")

	assert.Error(t, testError)
	assert.True(t, errors.Is(testError, os.ErrNotExist))
	assert.Equal(t, "could not open file foo: file does not exist", testError.Error())
}

func TestIs(t *testing.T) {
	testError := Wrap(os.ErrNotExist, "my error")

	assert.True(t, Is(testError, os.ErrNotExist))
}

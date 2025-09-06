package sorre

import (
	"errors"
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

type ErrTest struct {
	Code int
}

func (et *ErrTest) Error() string {
	return strconv.Itoa(et.Code)
}

func NewErrTest(code int) error {
	return &ErrTest{Code: code}
}

func WrapErr() error {
	err := errors.New("test error")
	return Wrap(err)
}

func ReturnIsError(err error) error {
	return Wrap(err)
}

func TestWrap(t *testing.T) {
	require.Equal(t, WrapErr().Error(), "[sorre.WrapErr:26] test error")
	require.Equal(t, Unwrap(WrapErr()).Error(), "test error")

	// Кастомная ошибка
	er403 := NewErrTest(403)
	var asErr *ErrTest
	require.True(t, errors.As(ReturnIsError(er403), &asErr))
	require.True(t, errors.Is(ReturnIsError(er403), er403))

	// Библиотечная ошибка
	er500 := errors.New("500")
	require.True(t, errors.Is(ReturnIsError(er500), er500))

	// Форматированная ошибка
	wrapErr := fmt.Errorf("Big error: %w", er403)
	require.True(t, errors.As(ReturnIsError(wrapErr), &asErr))
	require.True(t, errors.Is(ReturnIsError(wrapErr), er403))
	require.Equal(t, Unwrap(ReturnIsError(wrapErr)).Error(), "403")
}

func WrapfErr() error {
	err := errors.New("test error")
	return Wrapf(err, "message %d", 123)
}

func WrapfErrWithoutMsg() error {
	err := errors.New("test error")
	return Wrapf(err, "")
}

func WrapfErrWithoutArg() error {
	err := errors.New("test error")
	return Wrapf(err, "message")
}

func ReturnIsErrorf(err error) error {
	return Wrapf(err, "message %s", "")
}

func TestWrapf(t *testing.T) {
	require.Equal(t, WrapfErr().Error(), "[sorre.WrapfErr:56] message 123 : test error")
	require.Equal(t, WrapfErrWithoutMsg().Error(), "[sorre.WrapfErrWithoutMsg:61] : test error")
	require.Equal(t, WrapfErrWithoutArg().Error(), "[sorre.WrapfErrWithoutArg:66] message : test error")
	require.Equal(t, Unwrap(WrapfErr()).Error(), "test error")

	er403 := NewErrTest(403)
	var asErr *ErrTest
	require.True(t, errors.As(ReturnIsErrorf(er403), &asErr))
	require.True(t, errors.Is(ReturnIsErrorf(er403), er403))

	er500 := errors.New("500")
	require.True(t, errors.Is(ReturnIsErrorf(er500), er500))

	wrapErr := fmt.Errorf("Big error: %w", er403)
	require.True(t, errors.As(ReturnIsErrorf(wrapErr), &asErr))
	require.True(t, errors.Is(ReturnIsErrorf(wrapErr), er403))
	require.Equal(t, Unwrap(ReturnIsErrorf(wrapErr)).Error(), "403")
}

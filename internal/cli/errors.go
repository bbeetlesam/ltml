package cli

import (
	"errors"
	"fmt"
)

type ContextError struct {
	Err     error
	Context string
}

func (e *ContextError) Error() string {
	return fmt.Sprintf("%s: %s", e.Context, e.Err.Error())
}

var (
	ErrUsage          = errors.New("usage error")
	ErrInput          = errors.New("invalid input")
	ErrNotLTML        = errors.New("not an ltml file")
	ErrUnknownCommand = errors.New("unknown command")
	ErrNoFile         = errors.New("no file provided")
)

func FormatUserError(err *ContextError) string {
	switch err.Err {
	case ErrNoFile:
		return fmt.Sprintf("error: %s: no ltml file provided", err.Context)
	case ErrNotLTML:
		return fmt.Sprintf("error: %s: file must have .ltml extension", err.Context)
	case ErrUnknownCommand:
		return fmt.Sprintf("error: %s: unknown command\nForgot or lost? Try 'ltml help' for more information.", err.Context)
	default:
		return fmt.Sprintf("error: %s", err.Error())
	}
}

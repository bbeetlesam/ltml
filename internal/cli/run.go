package cli

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func Run(filename string) *ContextError {
	if filepath.Ext(filename) != ".ltml" {
		return &ContextError{Err: ErrNotLTML, Context: filename}
	}

	file, err := os.Open(filename)
	if err != nil {
		var pathErr *os.PathError
		if errors.As(err, &pathErr) {
			return &ContextError{Err: pathErr.Err, Context: filename}
		}
		return &ContextError{Err: err, Context: filename}
	}
	defer file.Close()

	fmt.Println("So now you have an LTML file. Nice. The Nice.")

	return nil
}

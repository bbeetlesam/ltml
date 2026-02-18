package cli

import (
	"fmt"
	"path/filepath"
)

func Run(filename string) *ContextError {
	if filepath.Ext(filename) != ".ltml" {
		return &ContextError{Err: ErrNotLTML, Context: filename}
	}

	fmt.Println("So now you have an LTML file. Nice. The Nice.")

	return nil
}

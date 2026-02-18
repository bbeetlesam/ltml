package cli

import (
	"fmt"
	"path/filepath"
)

func Build(filename string) *ContextError {
	if filepath.Ext(filename) != ".ltml" {
		return &ContextError{Err: ErrNotLTML, Context: filename}
	}

	fmt.Printf("can't build %s yet. it is not a master builder.", filename)
	return nil
}

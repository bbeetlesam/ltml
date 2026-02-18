package main

import (
	"fmt"
	"os"

	"github.com/bbeetlesam/ltml/internal/cli"
)

func main() {
	if err := cli.Execute(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, cli.FormatUserError(err))

		switch err.Err {
		case cli.ErrUsage, cli.ErrNoFile:
			os.Exit(2)
		case cli.ErrInput, cli.ErrNotLTML:
			os.Exit(3)
		case cli.ErrUnknownCommand:
			os.Exit(1)
		default:
			os.Exit(1)
		}
	}
}

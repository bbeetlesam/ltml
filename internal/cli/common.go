// Package cli implements the command-line interface for the LTML tool.
// It handles user-interface commands such as run, build, help, and version.
package cli

import (
	"fmt"
)

const (
	version     = "v0.0.0"
	versionName = "Preserved in Aspic"
)

const helpMsg = `CLI tool for compiling and running LTML/LTSS files.

USAGE:
  ltml <command> [options]

COMMANDS:
  run <file>      Run an LTML file
  build <file>    Build an LTML file to a plain shell script
  help            Show this very informational help
  version         Show the version of this program

EXAMPLES:
  $ ltml run file.ltml
  $ ltml build ready.ltml
  $ ltml help
`

func showHelp() {
	fmt.Print(helpMsg)
}

func showVersion() {
	fmt.Printf("%s - %s\n", version, versionName)
}

func Execute(args []string) *ContextError {
	if len(args) > 0 {
		command := args[0]

		var argv []string
		if len(args) > 1 {
			argv = args[1:]
		}

		switch command {
		case "run":
			if err := handleRun(argv); err != nil {
				return err
			}
		case "build":
			if err := handleBuild(argv); err != nil {
				return err
			}
		case "help":
			showHelp()
		case "version":
			showVersion()
		default:
			return &ContextError{Err: ErrUnknownCommand, Context: command}
		}
	} else {
		showHelp()
	}

	return nil
}

func handleRun(argv []string) *ContextError {
	if len(argv) == 0 {
		return &ContextError{Err: ErrNoFile, Context: "run"}
	}
	if err := Run(argv[0]); err != nil {
		return err
	}

	return nil
}

func handleBuild(argv []string) *ContextError {
	if len(argv) == 0 {
		return &ContextError{Err: ErrNoFile, Context: "build"}
	}
	if err := Build(argv[0]); err != nil {
		return err
	}

	return nil
}

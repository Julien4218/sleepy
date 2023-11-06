package main

import (
	"flag"
	"log"

	"github.com/Julien4218/sleepy/sleep"
	"github.com/spf13/cobra"
)

// Command represents the base command when called without any subcommands
var Command = &cobra.Command{
	PersistentPreRun:  globalInit,
	Use:               "sleepy",
	Short:             "A simple sleepy process",
	DisableAutoGenTag: true, // Do not print generation date on documentation
}

func init() {
	// Bind imported sub-commands
}

func globalInit(cmd *cobra.Command, args []string) {
}

func init() {
	Command.AddCommand(sleep.Command)
}

func main() {
	if err := Command.Execute(); err != nil {
		if err != flag.ErrHelp {
			log.Fatal(err)
		}
	}
}

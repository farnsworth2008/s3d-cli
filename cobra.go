package main

import "github.com/spf13/cobra"

// Declaring these here allows us to avoid importing `spf13/cobra` in other
// files.

type command = cobra.Command

var minArgs = cobra.MinimumNArgs

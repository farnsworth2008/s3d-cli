package main

import (
	"fmt"
	"strings"
)

func goCommands() *command {
	cGo.AddCommand(cGoSort)
	return cGo
}

var cGo = &command{
	Short: "Go Command Group",
	Use:   "go",

	Long: "This is the top level command for working with Go",
}

var cGoSort = &command{
	Short: "Sort Go files",
	Use:   "sort",

	Long: "Stub for a future command that will sort \"go\" files.",

	Run: func(cmd *command, args []string) {
		fmt.Println("Stub for sort go files: " + strings.Join(args, " "))
	},
}

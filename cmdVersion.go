// Copyright 2022 The S2D Club. All rights reserved.
// Use of this source code is governed by the LICENSE file.

package main

import "fmt"

// Format our version
func runVersion(cmd *command, args []string) {
	fmt.Println("0.0.1")
}

// Return our version command
func cmdVersion() *command {
	return versionCmd
}

// Our versionCmd command structure
var versionCmd = &command{
	Run:   runVersion,
	Short: "Display the program version",
	Use:   "version",
}

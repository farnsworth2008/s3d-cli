// Copyright 2022 The S2D Club. All rights reserved.
// Use of this source code is governed by the LICENSE file.

// The S3D Command Line Interface "main" package
package main

import "github.com/spf13/cobra"

// Run the CLI
func main() {
	s3dCmd.AddCommand(
		cmdAws(),
		cmdGo(),
		cmdTerraform(),
		cmdVersion(),
	)
	s3dCmd.Execute()
}

type command = cobra.Command

var s3dCmd = &command{
	Use:   "s3d",
	Short: "The CLI for the S3D Club",
	Long:  s3dLong,
}

var s3dLong = `
A collection of commands we find useful for software development work",
`

var minArgs = cobra.MinimumNArgs
var p1Int = 1
var p1String = ""

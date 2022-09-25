package main

// Run the CLI
func main() {
	s3dCmd.AddCommand(
		cmdAws(),
		newFlow(),
		cmdVersion(),
	)
	s3dCmd.Execute()
}

// We use the suffix `Cmd` for all command variables in the global scope
var s3dCmd = &command{
	Use:   "s3d",
	Short: "The CLI for the S3D Club",
	Long:  s3dLong,
}

var s3dLong = `
A collection of commands we find useful for software development work",
`

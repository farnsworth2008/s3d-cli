package main

// Command AWS creates the AWS command
func cmdAws() *command {
	s3dAws.AddCommand(cmdAwsEc2())
	return s3dAws
}

// The AWS command
var s3dAws = &command{
	Long:  "This is the top level command for working with AWS",
	Short: "AWS Command Group",
	Use:   "aws",
}

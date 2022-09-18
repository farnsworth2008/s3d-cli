package main

// Command AWS creates the AWS command
func cmdAws() *command {
	awsCmd.AddCommand(cmdAwsEc2())
	return awsCmd
}

// The AWS command
var awsCmd = &command{
	Long:  "This is the top level command for working with AWS",
	Short: "AWS Command Group",
	Use:   "aws",
}

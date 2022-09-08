package main

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var echoTimes int

func describeCheckGo() string {
	return "Check \"go\" files for style issues."
}

func describeCheckTerraform() string {
	return "Check \"tf\" files for style issues."
}

func describeEcho() string {
	return "Echo is for echoing anything back.\n" +
		"Echo works a lot like print, except it has a child command."
}

func describePrint() string {
	return "Print is for printing anything back to the screen.\n" +
		"For many years people have printed back to the screen."
}

func describeTimes() string {
	return "Echo things multiple times back to the user by providing\n" +
		"a count and a string."
}

func describeVersion() string {
	return "Show the version of S3D CLI"
}

func main() {
	cmdPrint := &cobra.Command{
		Args:  cobra.MinimumNArgs(1),
		Long:  describePrint(),
		Run:   runPrint,
		Short: "Demonstration command to print to the screen",
		Use:   "print [string to print]",
	}

	cmdEcho := &cobra.Command{
		Args:  cobra.MinimumNArgs(1),
		Long:  describeEcho(),
		Run:   runEcho,
		Short: "Demonstration command to echo to the screen",
		Use:   "echo [string to echo]",
	}

	cmdTimes := &cobra.Command{
		Args:  cobra.MinimumNArgs(1),
		Long:  describeTimes(),
		Run:   runTimes,
		Short: "Demonstration echoing content multiple times",
		Use:   "times [# times] [string to echo]",
	}

	cmdCheckGo := &cobra.Command{
		Args:  cobra.MinimumNArgs(1),
		Long:  describeCheckGo(),
		Run:   runCheckGo,
		Short: "Check Go style",
		Use:   "check-go [RegX]",
	}

	cmdCheckTerraform := &cobra.Command{
		Args:  cobra.MinimumNArgs(1),
		Long:  describeCheckTerraform(),
		Run:   runCheckTerraform,
		Short: "Check Terraform sytle",
		Use:   "check-tf [RegX]",
	}

	cmdVersion := &cobra.Command{
		Long:  describeVersion(),
		Run:   runVersion,
		Short: "Display the program version",
		Use:   "version",
	}

	echoTimes = 1
	rootCmd := &cobra.Command{Use: "s3d"}

	cmdTimes.Flags().IntVarP(
		&echoTimes, "times", "t", 1, "times to echo the input",
	)

	cmdEcho.AddCommand(cmdTimes)

	rootCmd.AddCommand(
		cmdPrint,
		cmdEcho,
		cmdCheckGo,
		cmdCheckTerraform,
		cmdVersion,
	)

	rootCmd.Execute()
}

func runCheckGo(cmd *cobra.Command, args []string) {
	fmt.Println("Stub for CheckGo: " + strings.Join(args, " "))
}

func runCheckTerraform(cmd *cobra.Command, args []string) {
	fmt.Println("Stub for CheckTerraform: " + strings.Join(args, " "))
}

func runEcho(cmd *cobra.Command, args []string) {
	fmt.Println("Echo: " + strings.Join(args, " "))
}

func runPrint(cmd *cobra.Command, args []string) {
	fmt.Println("Print: " + strings.Join(args, " "))
}

func runTimes(cmd *cobra.Command, args []string) {
	for i := 0; i < echoTimes; i++ {
		fmt.Println("Echo: " + strings.Join(args, " "))
	}
}

func runVersion(cmd *cobra.Command, args []string) {
	fmt.Println("Version: 0.0.0")
}

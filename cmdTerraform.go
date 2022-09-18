package main

import (
	"fmt"
	"strings"
)

func cmdTerraform() *command {
	terraform.AddCommand(terraformCheck)
	return terraform
}

func runTerraformCheck(cmd *command, args []string) {
	fmt.Println("Stub for check Terraorm: " + strings.Join(args, " "))
}

var terraform = &command{
	Long:  "This is the top level command for working with Terraform",
	Short: "Teraform Command Group",
	Use:   "tf",
}

var terraformCheck = &command{
	Args:  minArgs(1),
	Long:  "Stub for a future command that will check \"tf\" files.",
	Short: "Check Terraform style",
	Use:   "check [RegX]",
	Run:   runTerraformCheck,
}

package main

import "fmt"

// Command flow
func newFlowCommit() *command {
	cmdFlow.AddCommand(newFlowCommit())
	return cmdFlowCommit
}

// Create a new Flow command
func newFlow() *command {
	cmdFlow.AddCommand(cmdFlowCommit)
	return cmdFlow
}

// Flow commit command
func runFlowCommit(cmd *command, args []string) {
	fmt.Println("STUB: commit")
}

// The `flow` command
var cmdFlow = &command{
	Long:  longFlow,
	Short: "Flow Group",
	Use:   "flow",
}

// The `flow commmit` command
var cmdFlowCommit = &command{
	Long:  longFlowCommit,
	Short: "Make a flow commit",
	Use:   "commit",
	Run:   runFlowCommit,
}

var longFlow = `
Flow Commands are designed to help your software development work flow better.
Read about the flow commands at https://go.s3d.club/flow.
`

var longFlowCommit = `
The flow command command will stage and commit of all files from the work tree.

Prior to staging files content blocks of "CHANGES.md" will be sorted and the
pre-release version of the penultimate block in "CHANGES.md" will be
incremented.

The commit will be issued using text from the "CHANGES.md" file as shown in the
following example.

                ╔══════════════════════════════════════════════╗
                ║ 1.2.1-1 <task>; completed 3 of 5 items       ║
                ║                                              ║
                ║  - **TODO** Item one text                    ║
                ║  - **TODO** item two text                    ║
                ║  - **TODO** Item three text                  ║
                ║  - Completed item one                        ║
                ║  - Complete item two                         ║
                ║                                              ║
                ║  [S3D Flow Commit](https://go.s3d.club/flow) ║
                ╚══════════════════════════════════════════════╝

The value used for <task> will be the first value loaded from the ".s3d-task"
file in the current or any parent directory. If a ".s3d-task" file is not found
the task value "Changes" will be used.

Setting the "S3D_NO_AD" env var will suppress the markdown link.
`

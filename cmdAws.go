package main

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

// Command AWS creates the AWS command
func cmdAws() *command {
	awsCmd.AddCommand(cmdAwsEc2())
	return awsCmd
}

// AWS EC2 Command Goup godoc
// @Description The AWS EC2 Command Group provides commands that are useful
// when working with EC2.
// @Summary AWS EC2 Commands
// @Tags Thing
func cmdAwsEc2() *command {
	awsEc2Cmd.AddCommand(awsEc2StartCmd)
	awsEc2Cmd.AddCommand(awsEc2StopCmd)
	awsEc2Cmd.AddCommand(awsEc2R53Upsert)
	return awsEc2Cmd
}

// AWS EC2 Describe
func awsEc2Describe(
	todo context.Context,
	client *ec2.Client,
	name string,
	filterState string,
) *ec2.DescribeInstancesOutput {
	in := &ec2.DescribeInstancesInput{
		Filters: []types.Filter{{
			Name:   aws.String("instance-state-name"),
			Values: []string{filterState},
		}, {
			Name:   aws.String("tag:Name"),
			Values: []string{name},
		}},
	}
	in.InstanceIds = []string{}
	result, err := client.DescribeInstances(todo, in)
	kill(err)
	resLen := len(result.Reservations)
	killOn(0 == resLen, "ERROR: Did not find any instances with "+name)
	killOn(1 != resLen, "ERROR: Found more than one instance")
	return result
}

// Run EC2 Action
func runEc2Action(action string, name string) {
	todo := context.TODO()
	result := &ec2.DescribeInstancesOutput{}
	cfg, err := config.LoadDefaultConfig(todo)
	kill(err)

	running := []string{"running", "stopped"}
	stopped := []string{"stopped", "running"}
	actionMap := map[string][]string{
		"upsertR53": running,
		"start":     stopped,
		"stop":      running,
	}
	filterState := actionMap[action][0]
	nextFilterState := actionMap[action][1]

	client := ec2.NewFromConfig(cfg)
	result = awsEc2Describe(todo, client, name, filterState)
	inst := *result.Reservations[0].Instances[0].InstanceId

	switch action {
	case "start":
		startIn := &ec2.StartInstancesInput{}
		startIn.InstanceIds = []string{inst}
		_, err := client.StartInstances(todo, startIn)
		kill(err)
		time.Sleep(30 * time.Second)
		result = awsEc2Describe(todo, client, name, nextFilterState)
		dnsName := *result.Reservations[0].Instances[0].PublicDnsName
		applyR53CName(todo, cfg, name, dnsName)
	case "stop":
		stopIn := &ec2.StopInstancesInput{}
		stopIn.InstanceIds = []string{inst}
		_, err = client.StopInstances(todo, stopIn)
		kill(err)
	case "upsertR53":
		result = awsEc2Describe(todo, client, name, filterState)
		dnsName := *result.Reservations[0].Instances[0].PublicDnsName
		applyR53CName(todo, cfg, name, dnsName)
	default:
		fmt.Println("STUB: " + action + " " + inst)
	}
}

// Run EC2 Start
func runEc2Start(cmd *command, args []string) {
	runEc2Action("start", args[0])
}

// Run EC2 Upsert R53
func runEc2UpsertR53(cmd *command, args []string) {
	runEc2Action("upsertR53", args[0])
}

// AWS EC2 Command
var awsEc2Cmd = &command{
	Long:  "This is the top level command for working with AWS EC2",
	Short: "AWS EC2 Command Group",
	Use:   "ec2",
}

// AWS EC2 Start Command
var awsEc2StartCmd = &command{
	Args:  minArgs(1),
	Long:  longAwsEc2Start,
	Run:   runEc2Start,
	Short: "Start the AWS EC2 instance with a given name",
	Use:   "start [instance_name]",
}

// AWS EC2 Stop Command
var awsEc2StopCmd = &command{
	Args:  minArgs(1),
	Short: "Stop the AWS EC2 instance with a given name",
	Use:   "stop [instance_name]",
	Long:  longAwsEc2Stop,
	Run: func(cmd *command, args []string) {
		runEc2Action("stop", args[0])
	},
}

// AWS EC2 R53 Upsert Command
var awsEc2R53Upsert = &command{
	Args:  minArgs(1),
	Long:  longEc2UpsertR53,
	Run:   runEc2UpsertR53,
	Short: "Upserts a Route53 record for an EC2 instance",
	Use:   "r53-upsert instance-name",
}

// The AWS command
var awsCmd = &command{
	Long:  "This is the top level command for working with AWS",
	Short: "AWS Command Group",
	Use:   "aws",
}

var longAwsEc2Start = `
Starts an AWS EC2 instances using the supplied instance name to find the
instance.
`

var longAwsEc2Stop = `
Stops an AWS EC2 instances using the supplied instance name to find the
instance.
`

var longEc2UpsertR53 = `
Upserts a record into route53 for a given instance name
`

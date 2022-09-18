package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

// AWS EC2 Command Goup godoc
// @Description The AWS EC2 Command Group provides commands that are useful
// when working with EC2.
// @Summary AWS EC2 Commands
// @Tags Thing
func cmdAwsEc2() *command {
	awsEc2.AddCommand(awsEc2Start)
	awsEc2.AddCommand(awsEc2Stop)
	awsEc2.AddCommand(awsEc2UpsertR53)
	return awsEc2
}

var awsEc2 = &command{
	Long:  "This is the top level command for working with AWS EC2",
	Short: "AWS EC2 Command Group",
	Use:   "ec2",
}

var awsEc2Start = &command{
	Args:  minArgs(1),
	Long:  longAwsEc2Start,
	Run:   runEc2Start,
	Short: "Start the AWS EC2 instance with a given name",
	Use:   "start [instance_name]",
}

var awsEc2Stop = &command{
	Args:  minArgs(1),
	Short: "Stop the AWS EC2 instance with a given name",
	Use:   "stop [instance_name]",
	Long:  longAwsEc2Stop,
	Run: func(cmd *command, args []string) {
		runEc2Action("stop", args[0])
	},
}

var awsEc2UpsertR53 = &command{
	Args:  minArgs(1),
	Long:  longEc2UpsertR53,
	Run:   runEc2UpsertR53,
	Short: "Upserts a Route53 record for an EC2 instance",
	Use:   "upsert-r53 instance-name",
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

func awsEc2Describe(
	todo context.Context,
	client *ec2.Client,
	name string,
	filterState string,
) *ec2.DescribeInstancesOutput {
	in := &ec2.DescribeInstancesInput{
		Filters: []types.Filter{
			{
				Name:   aws.String("instance-state-name"),
				Values: []string{filterState},
			},
			{
				Name:   aws.String("tag:Name"),
				Values: []string{name},
			},
		},
	}
	in.InstanceIds = []string{}
	result, err := client.DescribeInstances(todo, in)
	if err != nil {
		log.Fatal(err)
	}

	if 0 == len(result.Reservations) {
		fmt.Println("ERROR: Did not find any instances with " + name)
	}

	if 1 != len(result.Reservations) {
		fmt.Println("ERROR: Found more than one instance")
	}

	return result
}

func runEc2Action(action string, name string) {
	todo := context.TODO()
	result := &ec2.DescribeInstancesOutput{}

	cfg, err := config.LoadDefaultConfig(todo)
	if err != nil {
		log.Fatal(err)
	}

	filterState := ""
	nextFilterState := ""
	switch action {
	case "upsertR53":
		filterState = "running"
	case "start":
		filterState = "stopped"
		nextFilterState = "running"
	case "stop":
		filterState = "running"
		nextFilterState = "stopped"
	default:
		panic(action)
	}

	client := ec2.NewFromConfig(cfg)
	result = awsEc2Describe(todo, client, name, filterState)
	inst := *result.Reservations[0].Instances[0].InstanceId

	switch action {
	case "start":
		startIn := &ec2.StartInstancesInput{}
		startIn.InstanceIds = []string{inst}
		_, err := client.StartInstances(todo, startIn)
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(30 * time.Second)
		result = awsEc2Describe(todo, client, name, nextFilterState)
		dnsName := *result.Reservations[0].Instances[0].PublicDnsName
		applyR53CName(todo, cfg, name, dnsName)
	case "stop":
		stopIn := &ec2.StopInstancesInput{}
		stopIn.InstanceIds = []string{inst}
		_, err = client.StopInstances(todo, stopIn)
		if err != nil {
			log.Fatal(err)
		}
	case "upsertR53":
		result = awsEc2Describe(todo, client, name, filterState)
		dnsName := *result.Reservations[0].Instances[0].PublicDnsName
		applyR53CName(todo, cfg, name, dnsName)
	default:
		fmt.Println("STUB: " + action + " " + inst)
	}
}

func runEc2Start(cmd *command, args []string) {
	runEc2Action("start", args[0])
}

func runEc2UpsertR53(cmd *command, args []string) {
	runEc2Action("upsertR53", args[0])
}

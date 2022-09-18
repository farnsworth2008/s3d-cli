package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
)

// The applyR53CName function applies a change to AWS and outputs the result.
func applyR53CName(
	context context.Context,
	cfg aws.Config,
	name string,
	dnsName string,
) {
	domain := os.Getenv("S3D_DOMAIN")
	hostedZone := os.Getenv("S3D_ZONE")
	name = name + "." + domain
	in := route53.ChangeResourceRecordSetsInput{
		ChangeBatch: &types.ChangeBatch{
			Changes: []types.Change{
				{
					Action: types.ChangeActionUpsert,
					ResourceRecordSet: &types.ResourceRecordSet{
						Name: &name,
						Type: types.RRTypeCname,
						TTL:  aws.Int64(60),
						ResourceRecords: []types.ResourceRecord{
							{
								Value: &dnsName,
							},
						},
					},
				},
			},
		},
		HostedZoneId: &hostedZone,
	}
	client := route53.NewFromConfig(cfg)
	result, err := client.ChangeResourceRecordSets(context, &in)
	if err != nil {
		log.Fatal(err)
	}
	encoding, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(encoding))
}

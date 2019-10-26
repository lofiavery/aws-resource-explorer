package aws_services

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	r53 "github.com/aws/aws-sdk-go/service/route53"
)

func GetZones(sess *session.Session) {
	r53 := r53.New(sess)
	r, err := r53.ListHostedZones(nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r.HostedZones)
	//https://docs.aws.amazon.com/sdk-for-go/api/service/route53/#Route53.ListResourceRecordSets
	// r53.ListResourceRecordSets()
}

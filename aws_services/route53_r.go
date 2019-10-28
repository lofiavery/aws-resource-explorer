package aws_services

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	route53 "github.com/aws/aws-sdk-go/service/route53"
)

func GetZones(sess *session.Session) ([]*route53.HostedZone, error) {
	r53 := route53.New(sess)
	r, err := r53.ListHostedZones(nil)
	if err != nil {
		return nil, err
	}
	return r.HostedZones, nil
}

func CreateListRecordSetInput(output *route53.ListResourceRecordSetsOutput, zoneId string) *route53.ListResourceRecordSetsInput {
	if output != nil {
		return &route53.ListResourceRecordSetsInput{
			HostedZoneId:          aws.String(zoneId), // Required
			StartRecordIdentifier: output.NextRecordIdentifier,
			StartRecordName:       output.NextRecordName,
			StartRecordType:       output.NextRecordType,
		}
	}
	return &route53.ListResourceRecordSetsInput{
		HostedZoneId: aws.String(zoneId),
	}
}

// GetRecordSetsNoPage gets all records of a given zone ignoring pagination - i.e getting all
func GetRecordSetsNoPage(sess *session.Session, zoneID string) ([]*route53.ResourceRecordSet, error) {
	result := []*route53.ResourceRecordSet{}
	r53 := route53.New(sess)
	var output *route53.ListResourceRecordSetsOutput = nil
	for {
		listParams := CreateListRecordSetInput(output, zoneID)
		response, err := r53.ListResourceRecordSets(listParams)
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
		output = response
		result = append(result, response.ResourceRecordSets...)
		if *response.IsTruncated == false {
			break
		}
	}
	return result, nil
}

/* Route53 zones */

type R53ZonesResource struct {
	Zones []*route53.HostedZone
}

func (i R53ZonesResource) Id() string {
	return "r53-zones"
}

type R53ZonesHandler struct {
	ZonesResource *R53ZonesResource
}

func (h R53ZonesHandler) Id() string {
	return "r53-zones"
}
func (h *R53ZonesHandler) Get() Resource {
	return h.ZonesResource
}
func (h *R53ZonesHandler) Fetch(config Conf, sess *session.Session, callback ResCb) {
	zones, err := GetZones(sess)
	if err != nil {
		callback(err, nil)
	} else {
		resource := &R53ZonesResource{Zones: zones}
		h.ZonesResource = resource
		callback(nil, resource)
	}
}

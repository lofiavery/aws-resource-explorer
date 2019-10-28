package aws_services

import "testing"
import (
	"aws_resource_explorer/utils"
)

func TestHostedZones(t *testing.T) {
	sess := utils.GetEnvSession()
	GetZones(sess)
}

func TestRecordSetsNoPage(t *testing.T) {
	sess := utils.GetEnvSession()
	res, _ := GetRecordSetsNoPage(sess, "/hostedzone/Z1AEL1R23MWQEZ")
	if len(res) < 250 {
		t.Errorf("TestEc2InstancesResource(): %d < 250", len(res))
	}
}

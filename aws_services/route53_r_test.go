package aws_services

import "testing"
import (
	"aws_resource_explorer/utils"
)

func TestHostedZones(t *testing.T) {
	sess := utils.GetEnvSession()
	GetZones(sess)
}



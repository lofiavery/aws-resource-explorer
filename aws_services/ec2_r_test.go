package aws_services

import "testing"
import (
	"aws_resource_explorer/utils"
)

func TestEc2InstancesResource(t *testing.T) {
	sess := utils.GetEnvSession()
	var conf = make(map[string]interface{})
	regions := []string{"us-east-1", "us-west-2"}
	conf["regions"] = regions
	ec2Handler := EC2Handler{}
	ec2Handler.Fetch(conf, sess, func(e error, r Resource) {
		if e != nil {
			t.Errorf("DescribeInstancesAWS() = %s", e)
		}
		ec2Instances := r.(*EC2Instances)
		if len(ec2Instances.Instances) < 500 {
			t.Errorf("TestEc2InstancesResource(): %d < 500", len(ec2Instances.Instances))
		}
	})
}

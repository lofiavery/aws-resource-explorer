package main

import (
	s "aws_resource_explorer/aws_services"
	"aws_resource_explorer/utils"
	"fmt"
)

func main() {
	sess := utils.GetEnvSession()
	var conf = make(map[string]interface{})
	regions := []string{"us-east-1", "us-west-2"}
	conf["regions"] = regions
	ec2Handler := s.EC2Handler{}
	ec2Handler.Fetch(conf, sess, func(e error, r s.Resource) {
		if e != nil {
			panic(e)
		}
		ec2Instances := r.(*s.EC2Instances)
		fmt.Println(len(ec2Instances.Instances))
	})
}

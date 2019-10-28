package main

import (
	s "aws_resource_explorer/aws_services"
	"aws_resource_explorer/utils"
	"fmt"
)

func main() {
	fmt.Println("...")
	sess := utils.GetEnvSession()
	// var conf = make(map[string]interface{})
	// regions := []string{"us-east-1", "us-west-2"}
	// conf["regions"] = regions
	// ec2Handler := s.EC2InstanceHandler{}
	// ec2Handler.Fetch(conf, sess, func(e error, r s.Resource) {
	// 	if e != nil {
	// 		panic(e)
	// 	}
	// 	ec2Instances := r.(*s.EC2InstancesResource)
	// 	fmt.Println(ec2Instances.Instances)
	// })

	// enis
	// var conf = make(map[string]interface{})
	// regions := []string{"us-east-1", "us-west-2"}
	// conf["regions"] = regions
	// eniHandler := s.EC2EniHander{}
	// eniHandler.Fetch(conf, sess, func(e error, r s.Resource) {
	// 	if e != nil {
	// 		panic(e)
	// 	}
	// 	enis := r.(*s.EC2EniResource)
	// 	fmt.Println(len(enis.Interfaces))
	// })

	// route53
	res, _ := s.GetRecordSetsNoPage(sess, "/hostedzone/Z1AEL1R23MWQEZ")
	fmt.Println(len(res))
}

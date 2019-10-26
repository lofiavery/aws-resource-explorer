package main

import (
	s "aws_resource_explorer/aws_services"
	"aws_resource_explorer/utils"
	"fmt"
)

func main() {
	sess := utils.GetEnvSession()
	instances, _ := s.GetInstances(sess)
	fmt.Println(len(instances))
	// s.DescribeInstancesAWS2()
	fmt.Println(s.Hello2())
}

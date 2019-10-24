package main

import (
	s "aws_resource_explorer/aws_services"
	"fmt"
)

func main() {
	s.DescribeInstancesAWS2()
	fmt.Println(s.Hello2())
}

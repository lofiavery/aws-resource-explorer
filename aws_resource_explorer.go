package aws_resource_explorer

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	ec2 "github.com/aws/aws-sdk-go/service/ec2"
)

func Hello() string {
	return "Hello, world."
}

func DescribeInstancesAWS() bool {
	// Load session from shared config
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create new EC2 client
	ec2Svc := ec2.New(sess)
	// Call to get detailed information on each instance
	result, err := ec2Svc.DescribeInstances(nil)
	if err != nil {
		fmt.Println("Error", err)
	} else {
		// fmt.Println("Success", result)
		reservations := result.Reservations
		counter := 0
		fmt.Println(len(reservations))
		for _, r := range reservations {
			for _, instance := range r.Instances {
				// fmt.Println(*instance.InstanceId)
				fmt.Println(*instance.State.Name)
				counter++
			}
		}
		fmt.Printf("Total of %v instances \n", counter)
		mapping := MapInstanceStates(reservations)
		fmt.Println(mapping)

	}
	return true
}

func MapInstanceStates(reservations []*ec2.Reservation) map[string]int {
	m := make(map[string]int)
	for _, r := range reservations {
		for _, instance := range r.Instances {
			state := *instance.State.Name
			m[state]++
		}
	}
	return m
}

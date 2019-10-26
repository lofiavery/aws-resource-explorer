package aws_services

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	ec2 "github.com/aws/aws-sdk-go/service/ec2"
)

func Cool() {
	fmt.Println("Cool is running! ")
}
func Hello2() string {
	return "Hello, world."
}

func DescribeInstancesAWS2() bool {
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
		mapping := MapInstanceStates2(reservations)
		fmt.Println(mapping)
		flatInstances := FlatReservations2(reservations)
		fmt.Println("Len of flat instances: ", len(flatInstances))

	}
	return true
}

func MapInstanceStates2(reservations []*ec2.Reservation) map[string]int {
	m := make(map[string]int)
	for _, r := range reservations {
		for _, instance := range r.Instances {
			state := *instance.State.Name
			m[state]++
		}
	}
	return m
}

func FlatReservations2(reservations []*ec2.Reservation) []*ec2.Instance {
	instances := []*ec2.Instance{}
	for _, r := range reservations {
		for _, i := range r.Instances {
			instances = append(instances, i)
		}
	}
	return instances
}

// Get all ec2 instances running in the account
func GetInstances(sess *session.Session) ([]*ec2.Instance, error) {
	eastInstances, err1 := GetInstancesRegion(sess, "us-east-1")
	if err1 != nil {
		return nil, err1
	}
	westInstances, err2 := GetInstancesRegion(sess, "us-west-2")
	if err2 != nil {
		return nil, err2
	}
	return append(eastInstances, westInstances...), nil
}

func GetInstancesRegion(sess *session.Session, region string) ([]*ec2.Instance, error) {
	sess.Config.Region = aws.String(region)
	ec2Svc := ec2.New(sess)
	result, err := ec2Svc.DescribeInstances(nil)
	if err != nil {
		return nil, err
	}
	instances := FlatReservations2(result.Reservations)
	return instances, nil
}
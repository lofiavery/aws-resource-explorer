package aws_services

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	ec2 "github.com/aws/aws-sdk-go/service/ec2"
	"sync"
)

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
func GetEniRegion(sess *session.Session, region string) ([]*ec2.NetworkInterface, error) {
	sess.Config.Region = aws.String(region)
	ec2Svc := ec2.New(sess)
	result, err := ec2Svc.DescribeNetworkInterfaces(nil)
	if err != nil {
		return nil, err
	}
	return result.NetworkInterfaces, nil
}

/*
	Network Interfaces (eni)
*/
type EC2EniResource struct {
	Interfaces []*ec2.NetworkInterface
}

func (i EC2EniResource) Id() string {
	return "ec2-network-interfaces"
}

type EC2EniHander struct {
	Eni *EC2EniResource
}

func (h EC2EniHander) Id() string {
	return "ec2-network-interfaces"
}
func (h *EC2EniHander) Get() Resource {
	return h.Eni
}
func (h *EC2EniHander) Fetch(config Conf, sess *session.Session, callback ResCb) {
	regions := []string{}
	if config["regions"] != nil {
		regions = config["regions"].([]string)
	} else {
		regions = []string{*sess.Config.Region}
	}
	var wg sync.WaitGroup
	wg.Add(len(regions))
	allEnis := []*ec2.NetworkInterface{}
	for _, r := range regions {
		go func(r string) {
			defer wg.Done()
			enis, err := GetEniRegion(sess, r)
			if err == nil {
				allEnis = append(allEnis, enis...)
			}
		}(r)
	}
	wg.Wait()
	resource := &EC2EniResource{Interfaces: allEnis}
	h.Eni = resource
	callback(nil, resource)
}

/*
	Instances
*/

type EC2InstancesResource struct {
	Instances []*ec2.Instance
}

func (i EC2InstancesResource) Id() string {
	return "ec2-instances"
}

type EC2InstanceHandler struct {
	instances *EC2InstancesResource
}

func (h EC2InstanceHandler) Id() string {
	return "ec2-instances"
}
func (h *EC2InstanceHandler) Get() Resource {
	return h.instances
}
func (h *EC2InstanceHandler) Fetch(config Conf, sess *session.Session, callback ResCb) {
	regions := []string{}
	if config["regions"] != nil {
		regions = config["regions"].([]string)
	} else {
		regions = []string{*sess.Config.Region}
	}
	var wg sync.WaitGroup
	wg.Add(len(regions))
	allInstances := []*ec2.Instance{}
	for _, r := range regions {
		go func(r string) {
			defer wg.Done()
			instances, err := GetInstancesRegion(sess, r)
			if err == nil {
				allInstances = append(allInstances, instances...)
			}
		}(r)
	}
	wg.Wait()
	resource := &EC2InstancesResource{Instances: allInstances}
	h.instances = resource
	callback(nil, resource)
}

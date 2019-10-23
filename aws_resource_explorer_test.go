package aws_resource_explorer

import "testing"

func TestHello(t *testing.T) {
	want := "Hello, world."
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}

func TestDescribeInstancesAWS(t *testing.T) {
	want := true
	if success := DescribeInstancesAWS(); success != want {
		t.Errorf("DescribeInstancesAWS() = %t, want %t", success, want)
	}
}

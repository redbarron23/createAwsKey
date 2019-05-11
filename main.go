package main

// delete if afterwards
// aws ec2 delete-key-pair --key-name test2

import (
	"fmt"
)

var pairName string

func init() {
	pairName = InitKey()
}

// InitKey persist keyname throughout
func InitKey() string {
	return "test2"
}

func main() {
	fmt.Println("Demonstrating creating and deleting AWS Key")

	createAwsKey()
	deleteAwsKey()
}

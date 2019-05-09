package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// Creates a new EC2 key pair for the name provided.
//
// Usage:
//    go run ec2_create_keypair.go KEY_PAIR_NAME
func main() {
	if len(os.Args) != 2 {
		exitErrorf("pair name required\nUsage: %s key_pair_name",
			filepath.Base(os.Args[0]))
	}
	pairName := os.Args[1]
	fileName := pairName + ".pem"

	// Initialize a session in eu-west-2 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-west-2")},
	)

	// Create an EC2 service client.
	svc := ec2.New(sess)

	// Creates a new  key pair with the given name
	result, err := svc.CreateKeyPair(&ec2.CreateKeyPairInput{
		KeyName: aws.String(pairName),
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok && aerr.Code() == "InvalidKeyPair.Duplicate" {
			exitErrorf("Keypair %q already exists.", pairName)
		}
		exitErrorf("Unable to create key pair: %s, %v.", pairName, err)
	}

	// lets write to a pem file and STDOUT
	f, err := os.Create(fileName)
	check(err)
	defer f.Close()

	fmt.Printf("Created key pair %q %s\n",
		*result.KeyName, *result.KeyFingerprint)

	w := bufio.NewWriter(f)
	check(err)
	w.Flush()

}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

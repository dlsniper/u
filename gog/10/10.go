package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func consumeSQS() {

	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := sqs.New(sess)

	params := &sqs.ReceiveMessageInput{
		QueueUrl: aws.String("https://sqs.eu-west-1.amazonaws.com/939326250094/cc3-shs01-inf-awsconfig-logbucket-FileCreated"), // Required
		MaxNumberOfMessages: aws.Int64(1),
		VisibilityTimeout:       aws.Int64(30),
		WaitTimeSeconds:         aws.Int64(15),
	}
	resp, err := svc.ReceiveMessage(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)

}

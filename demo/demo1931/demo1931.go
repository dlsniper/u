package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sns"
)

func main() {
	abc := 123
	fmt.Println(abc)
	config := aws.Config{Region: aws.String("us-west-2")}
	svc := sns.New(&config)
	fmt.Println(svc.APIVersion)
}

package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rds"
)

type MyEvent struct {
	InstanceRegion string   `json:"InstanceRegion"`
	InstanceIDList []string `json:"InstanceIdList"`
}

func HandleLambdaEvent(event MyEvent) (string, error) {
	// instances id list
	for _, id := range event.InstanceIDList {
		input := &rds.StartDBInstanceInput{
			DBInstanceIdentifier: aws.String(id),
		}

		// start instances
		svc := rds.New(session.New())
		result, err := svc.StartDBInstance(input)
		if err == nil {
			fmt.Println(result)

		} else {
			fmt.Println(err.Error())
		}
	}

	return fmt.Sprintf("Start DB instances %d!", len(event.InstanceIDList)), nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}

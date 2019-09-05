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
		//input.InstanceIds = append(input.InstanceIds, aws.String(id))
		input := &rds.StopDBInstanceInput{
			DBInstanceIdentifier: aws.String(id),
		}

		// stop instances
		svc := rds.New(session.New())
		result, err := svc.StopDBInstance(input)
		if err == nil {
			fmt.Println(result)

		} else {
			fmt.Println(err.Error())
		}
	}

	return fmt.Sprintf("Stop instances %d!", len(event.InstanceIDList)), nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}

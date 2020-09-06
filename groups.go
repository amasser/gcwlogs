package cwlogs

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

// CreateLogGroup is a function creates a log group.
func (t CWLogs) CreateLogGroup(groupName string) error {
	var createLogGroupIn *cloudwatchlogs.CreateLogGroupInput = &cloudwatchlogs.CreateLogGroupInput{
		LogGroupName: aws.String(groupName),
	}

	_, err := t.client.CreateLogGroup(createLogGroupIn)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}

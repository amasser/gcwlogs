package cwlogs

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

// CreateLogStream is a function creates a log stream.
func (t CWLogs) CreateLogStream(groupName string, streamName string) error {
	var createLogStreamIn *cloudwatchlogs.CreateLogStreamInput = &cloudwatchlogs.CreateLogStreamInput{
		LogGroupName:  aws.String(groupName),
		LogStreamName: aws.String(streamName),
	}

	_, err := t.client.CreateLogStream(createLogStreamIn)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}

// GetNextSequenceToken is function gets the NextSequenceToken of a log stream.
func (t CWLogs) GetNextSequenceToken(groupName string, streamName string) *string {
	var nextSeqToken string = ""

	var describeLogStreamsIn *cloudwatchlogs.DescribeLogStreamsInput = &cloudwatchlogs.DescribeLogStreamsInput{
		LogStreamNamePrefix: aws.String(streamName),
		LogGroupName:        aws.String(groupName),
	}
	out, err := t.client.DescribeLogStreams(describeLogStreamsIn)

	if err != nil {
		fmt.Println(err.Error())
		return aws.String("")
	}

	if len(out.LogStreams) == 0 {
		return aws.String("")
	}

	for _, logStream := range out.LogStreams {
		if *logStream.LogStreamName == streamName {
			if logStream.UploadSequenceToken != nil {
				nextSeqToken = *logStream.UploadSequenceToken
			}
		}
	}

	return aws.String(nextSeqToken)
}

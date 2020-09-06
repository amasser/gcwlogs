package cwlogs

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

// CreateLogEvent is a function creates a log event for PutLogEvents.
func (t CWLogs) CreateLogEvent(message string, timestamp int64) (*cloudwatchlogs.InputLogEvent, error) {
	if message == "" {
		errMessage := "Message for a log event is required."
		return nil, errors.New(errMessage)
	}

	if timestamp <= 0 {
		errMessage := "Invalid timestamp for a log event."
		return nil, errors.New(errMessage)
	}

	timestampDigits := len(strconv.FormatInt(timestamp, 10))
	var multipyNum int64 = 10
	for n := 0; n < (13 - timestampDigits); n++ {
		timestamp = timestamp * multipyNum
	}

	logEvent := &cloudwatchlogs.InputLogEvent{
		Message:   aws.String(message),
		Timestamp: aws.Int64(timestamp),
	}

	return logEvent, nil
}

// PutLogEvents is a function puts log event
func (t CWLogs) PutLogEvents(groupName string, streamName string, logEvents []*cloudwatchlogs.InputLogEvent) error {
	var putLogEventIn *cloudwatchlogs.PutLogEventsInput = &cloudwatchlogs.PutLogEventsInput{
		LogGroupName:  aws.String(groupName),
		LogStreamName: aws.String(streamName),
		LogEvents:     logEvents,
	}

	var nextSeqToken *string = t.GetNextSequenceToken(groupName, streamName)
	if *nextSeqToken != "" {
		putLogEventIn.SequenceToken = nextSeqToken
	}

	_, err := t.client.PutLogEvents(putLogEventIn)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}

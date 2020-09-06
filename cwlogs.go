package cwlogs

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

// CWLogs is main struct.
type CWLogs struct {
	client *cloudwatchlogs.CloudWatchLogs
}

// New is a construct function.
func New(region string) *CWLogs {
	cwlogs := &CWLogs{
		client: cloudwatchlogs.New(
			session.Must(session.NewSession()),
			aws.NewConfig().WithRegion(region)),
	}

	return cwlogs
}

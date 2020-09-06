go-cwlogs
===

go-cwlogs is a simple client for putting log events to CloudWatch Logs from applications implemented with Golang.

## Install

```bash
go get michimani/go-cwlogs
```

## Example

```go
package main

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"

	gocwlogs "github.com/michimani/go-cwlogs"
)

func main() {
	cwl := gocwlogs.New("ap-northeast-1")

	var groupName string = "smaple/log/group"
	var streamName string = "sample/log/stream"

	cwl.CreateLogGroup(groupName)
	cwl.CreateLogStream(groupName, streamName)

	var eventMessage string = "[%d] This is a sample message for a log event."
	var eventTimestamp int64 = time.Now().Unix()
	logEvent1, _ := cwl.CreateLogEvent(fmt.Sprintf(eventMessage, 1), eventTimestamp)
	logEvent2, _ := cwl.CreateLogEvent(fmt.Sprintf(eventMessage, 2), eventTimestamp)

	var logEvents []*cloudwatchlogs.InputLogEvent = []*cloudwatchlogs.InputLogEvent{
		logEvent1,
		logEvent2,
	}

	cwl.PutLogEvents(groupName, streamName, logEvents)
}
```
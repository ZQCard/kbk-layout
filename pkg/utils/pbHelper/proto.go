package pbHelper

import (
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
)

// Convert a protobuf Timestamp to a Golang time.Time object.
func ProtoTimestampToTime(ts *timestamp.Timestamp) time.Time {
	t, _ := ptypes.Timestamp(ts)
	return t
}

func TimeToProtoTimestamp(t time.Time) (ts *timestamp.Timestamp) {
	ts, _ = ptypes.TimestampProto(t)
	return ts
}

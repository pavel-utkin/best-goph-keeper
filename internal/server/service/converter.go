package service

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"time"
)

func ConvertTimeToTimestamp(t time.Time) (*timestamp.Timestamp, error) {
	ts, err := ptypes.TimestampProto(t)
	if err != nil {
		return nil, err
	}
	return ts, nil
}

func ConvertTimestampToTime(ts *timestamp.Timestamp) (time.Time, error) {
	t, err := ptypes.Timestamp(ts)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

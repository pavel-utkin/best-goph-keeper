package service

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func ConvertTimeToTimestamp(t time.Time) (*timestamp.Timestamp, error) {
	return timestamppb.New(t), nil
}

func ConvertTimestampToTime(ts *timestamp.Timestamp) (time.Time, error) {
	err := ts.CheckValid()
	if err != nil {
		return time.Time{}, err
	}
	return ts.AsTime(), nil
}

package types

import (
	"log"
	"time"

	"github.com/yash91989201/superfast-delivery-api/common/pb"
	"go.mongodb.org/mongo-driver/v2/bson"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Date struct {
	Year  int32
	Month int32
	Day   int32
}

func PtrToString(s string) *string    { return &s }
func PtrToBool(b bool) *bool          { return &b }
func PtrToFloat64(f float64) *float64 { return &f }

func (d *Date) ToTime() *time.Time {
	t := time.Date(int(d.Year), time.Month(d.Month), int(d.Day), 0, 0, 0, 0, time.UTC)
	return &t
}

func ToBoolPtr(b bool) *bool {
	return &b
}

func ToStrPtr(s string) *string {
	return &s
}

func HexToObjectID(id string) bson.ObjectID {
	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("Invalid object id: %v", err)
	}

	return objectID
}

func ToPbTimestamp(t time.Time) *timestamppb.Timestamp {
	if t.IsZero() {
		return nil
	}
	return timestamppb.New(t)
}

func ToTime(ts *timestamppb.Timestamp) time.Time {
	if ts != nil {
		return ts.AsTime()
	}
	return time.Time{}
}

func ToTimePtr(ts *timestamppb.Timestamp) *time.Time {
	if ts != nil {
		t := ts.AsTime()
		return &t
	}
	return nil
}

func ToDate(d *pb.Date) *Date {
	return &Date{
		Year:  d.Year,
		Month: d.Month,
		Day:   d.Day,
	}
}

func ToPbDate(d *Date) *pb.Date {
	return &pb.Date{
		Year:  d.Year,
		Month: d.Month,
		Day:   d.Day,
	}
}

func TimeToPbDate(t *time.Time) *pb.Date {
	if t == nil {
		return nil
	}

	return &pb.Date{
		Year:  int32(t.Year()),
		Month: int32(t.Month()),
		Day:   int32(t.Day()),
	}
}

func PbTimeStampToStrPtr(t *timestamppb.Timestamp) *string {
	if t == nil {
		return nil
	}

	tm := t.AsTime()

	// Format as ISO 8601 in UTC
	timeStr := tm.UTC().Format(time.RFC3339)
	return &timeStr
}

func PbDateToTime(d *pb.Date) *time.Time {
	if d == nil {
		return nil
	}
	t := time.Date(
		int(d.Year),
		time.Month(d.Month),
		int(d.Day),
		0, 0, 0, 0,
		time.UTC,
	)
	return &t
}

func TimePtrToPbTime(t *time.Time) *timestamppb.Timestamp {
	if t == nil {
		return nil
	}

	return timestamppb.New(*t)
}

func stringsToBsonObjectIDs(ids []string) []bson.ObjectID {
	objectIDs := make([]bson.ObjectID, len(ids))
	for i, id := range ids {
		objectIDs[i] = HexToObjectID(id)
	}
	return objectIDs
}

func bsonObjectIDsToStrings(ids []bson.ObjectID) []string {
	strIDs := make([]string, len(ids))
	for i, id := range ids {
		strIDs[i] = id.Hex()
	}
	return strIDs
}

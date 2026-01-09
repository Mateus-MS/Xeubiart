package internal_datetime

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

var (
	ErrNonUTCTime = errors.New("in order to create a UTCTime, the given timezone should be setted as UTC")
)

// UTCTime wraps time.Time to represent a UTC timestamp
type UTCTime struct {
	time.Time
}

// NewUTCTimeFromTime creates a UTCTime from a time that must already be in UTC.
func NewUTCTimeFromTime(t time.Time) (*UTCTime, error) {
	if t.Location() != time.UTC {
		return nil, ErrNonUTCTime
	}

	return &UTCTime{
		Time: t.UTC(),
	}, nil
}

// Converts the given UTCTime to a given timezone
func (ut *UTCTime) ToLocalTime(local *time.Location) (*LocalTime, error) {
	if local != time.UTC {
		return nil, ErrNonUTCTime
	}

	return &LocalTime{
		Time: ut.Time.In(local),
	}, nil
}

// Custom BSON marshaling ensures Mongo stores it as a proper date,
// so range queries ($gte / $lt) work correctly.
func (t UTCTime) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bson.MarshalValue(t.Time)
}

// Custom BSON unmarshaling ensures Mongo reads it as a proper date,
// so range queries ($gte / $lt) work correctly.
func (t *UTCTime) UnmarshalBSONValue(bt bsontype.Type, data []byte) error {
	return bson.UnmarshalValue(bt, data, &t.Time)
}

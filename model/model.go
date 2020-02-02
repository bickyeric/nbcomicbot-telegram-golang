package model

import (
	"errors"
	"io"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MarshalTimestamp ...
func MarshalTimestamp(t time.Time) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.FormatInt(t.Unix(), 10))
	})
}

// UnmarshalTimestamp ...
func UnmarshalTimestamp(v interface{}) (time.Time, error) {
	if tmpStr, ok := v.(int64); ok {
		return time.Unix(tmpStr, 0), nil
	}
	return time.Time{}, errors.New("time should be a unix timestamp")
}

// MarshalID ...
func MarshalID(v primitive.ObjectID) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.Quote(v.Hex()))
	})
}

// UnmarshalID ...
func UnmarshalID(v interface{}) (primitive.ObjectID, error) {
	if tmpStr, ok := v.(string); ok {
		return primitive.ObjectIDFromHex(tmpStr)
	}
	return primitive.NewObjectID(), errors.New("id is not valid object id")
}

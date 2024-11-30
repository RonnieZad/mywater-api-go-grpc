package utils

import wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"

func WrapString(str string) *wrapperspb.StringValue {
    return &wrapperspb.StringValue{
        Value: str,
    }
}

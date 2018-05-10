package types

import (
	"fmt"
	"time"
)

func Int(v int64) *int64 {
	return &v
}

func Float(v float64) *float64 {
	return &v
}

func Bool(v bool) *bool {
	return &v
}

func String(v string) *string {
	sv := fmt.Sprintf("%v", v)
	return &sv
}

func Time(v time.Time) *time.Time {
	return &v
}

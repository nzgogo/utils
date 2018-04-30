package types

import "time"

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
	return &v
}

func Time(v time.Time) *time.Time {
	return &v
}

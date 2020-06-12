package lib

import (
	"aging/methods"
	"time"
)

const (
	LayoutISO = "2006-01-02"
	LayoutUS  = "January 2, 2006"
	Custom    = "1/2/2006 0:00:00"
	Batch     = "2006-1"
)

func GetTime(t time.Time, err error) time.Time {
	methods.FailOnError(err)
	return t
}

func GetFloat(d float64, err error) float64 {
	methods.FailOnError(err)
	return d
}

func GetInt(d float64, err error) int64 {
	methods.FailOnError(err)
	return int64(d * 1_000)
}

package model

import "time"

type Date time.Time

func NewDate(t time.Time) Date {
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	return Date(time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, jst))
}

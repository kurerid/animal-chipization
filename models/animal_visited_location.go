package models

import "time"

type AnimalVisitedLocation struct {
	Id                           int64
	DateTimeOfVisitLocationPoint time.Time
	LocationPointId              int64
}

package models

import "time"

type Animal struct {
	Id                 int64
	AnimalTypes        []int64
	Weight             float32
	Length             float32
	Height             float32
	Gender             string
	LifeStatus         string
	ChippingDateTime   time.Time
	ChipperId          int
	ChippingLocationId int64
	VisitedLocations   int64
	DeathDateTime      time.Time
}

package models

type AnimalVisitedLocation struct {
	Id                           int64  `json:"id"`
	DateTimeOfVisitLocationPoint string `json:"dateTimeOfVisitLocationPoint"`
	LocationPointId              int64  `json:"locationPointId"`
}

type AnimalVisitedLocationListInput struct {
	Animal    int64  `form:"animalId" binding:"required"`
	StartDate string `form:"startDateTime" binding:"required"`
	EndDate   string `form:"endDateTime" binding:"required"`
	Offset    int    `form:"from" binding:"required"`
	Limit     int    `form:"size" binding:"required"`
}

type AnimalVisitedLocationNewInput struct {
	Animal   int64 `form:"animalId" binding:"required"`
	Location int64 `form:"pointId" binding:"required"`
}

type AnimalVisitedLocationChangeInput struct {
	Animal          int64 `form:"animalId" binding:"required"`
	VisitedLocation int64 `json:"visitedLocationPointId" binding:"required"`
	Location        int64 `json:"locationPointId" binding:"required"`
}

type AnimalVisitedLocationRemoveInput struct {
	Animal          int64 `form:"animalId"`
	VisitedLocation int64 `form:"visitedPointId"`
}

package models

type Location struct {
	Id        int64   `json:"id"`
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

type LocationGetByIdInput struct {
	Id int64 `form:"pointId" binding:"required,numeric"`
}

type LocationNewInput struct {
	Latitude  float32 `json:"latitude" binding:"required"`
	Longitude float32 `json:"longitude" binding:"required"`
}

type LocationModifyInput struct {
	Id        int64   `form:"pointId" binding:"required,numeric"`
	Latitude  float32 `json:"latitude" binding:"required"`
	Longitude float32 `json:"longitude" binding:"required"`
}

type LocationRemoveInput struct {
	Id int64 `form:"pointId" binding:"required,numeric"`
}

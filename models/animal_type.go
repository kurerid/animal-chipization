package models

type AnimalType struct {
	Id   int64
	Type string
}

type AnimalTypeGetByIdInput struct {
	Id int64 `form:"typeId" binding:"required,numeric"`
}

type AnimalTypeNewInput struct {
	Type string `json:"type" binding:"required"`
}

type AnimalTypeModifyInput struct {
	Id   int64  `form:"typeId" binding:"required,numeric"`
	Type string `json:"type" binding:"required"`
}

type AnimalTypeRemoveInput struct {
	Id int64 `form:"typeId"`
}

package models

type Animal struct {
	Id                 int64
	AnimalTypes        []int64
	Weight             float32
	Length             float32
	Height             float32
	Gender             string
	LifeStatus         string
	ChippingDateTime   string
	ChipperId          int
	ChippingLocationId int64
	VisitedLocations   int64
	DeathDateTime      string
}

type AnimalGetByIdInput struct {
	Id int64 `form:"animalId" binding:"required,numeric"`
}

type AnimalSearchInput struct {
	StartDate        string `form:"startDateTime"`
	EndDate          string `form:"endDateTime"`
	Chipper          int    `form:"chipperId"`
	ChippingLocation int64  `form:"chippingLocationId"`
	LifeStatus       string `form:"lifeStatus"`
	Gender           string `form:"gender"`
	Offset           int    `form:"from"`
	Limit            int    `form:"size"`
}

type AnimalNewInput struct {
	Types            []int64 `json:"animalTypes" binding:"required"`
	Weight           float32 `json:"weight" binding:"required"`
	Length           float32 `json:"length" binding:"required"`
	Height           float32 `json:"height" binding:"required"`
	Gender           string  `json:"gender" binding:"required"`
	Chipper          int     `json:"chipperId" binding:"required"`
	ChippingLocation int64   `json:"chippingLocationId" binding:"required"`
}

type AnimalUpdateInput struct {
	Id               int64   `form:"animalId" binding:"required"`
	Weight           float32 `json:"weight" binding:"required"`
	Length           float32 `json:"length" binding:"required"`
	Height           float32 `json:"height" binding:"required"`
	Gender           string  `json:"gender" binding:"required"`
	LifeStatus       string  `json:"lifeStatus" binding:"required"`
	Chipper          int     `json:"chipperId" binding:"required"`
	ChippingLocation int64   `json:"chippingLocationId" binding:"required"`
}

type AnimalRemoveInput struct {
	Id int64 `form:"animalId" binding:"required,numeric"`
}

type AnimalAddTypeInput struct {
	Animal int64 `form:"animalId" binding:"required,numeric"`
	Type   int64 `form:"typeId" binding:"required,numeric"`
}

type AnimalChangeTypeInput struct {
	Animal  int64 `form:"animalId" binding:"required"`
	OldType int64 `json:"oldTypeId" binding:"required"`
	NewType int64 `json:"newTypeId" binding:"required"`
}

type AnimalRemoveTypeInput struct {
	Animal int64 `form:"animalId" binding:"required,numeric"`
	Type   int64 `form:"typeId" binding:"required,numeric"`
}

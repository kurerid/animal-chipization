package service

import (
	"animal-chipization/models"
	"animal-chipization/pkg/repository"
)

type Authorization interface {
	SignUp(input *models.SignUpInput) (*models.SignUpOutput, error)
	SignIn()
}

type User interface {
	GetById(id int) (*models.AccountGetByIdOutput, error)
	Search(input *models.AccountSearchInput) (*models.AccountSearchOutput, error)
	Update(input *models.AccountUpdateInput) (*models.AccountUpdateOutput, error)
	Remove(id int) error
}

type Animal interface {
	GetById(id int64) (*models.Animal, error)
	Search(input *models.AnimalSearchInput) (*Animal, error)
	New(input *models.AnimalNewInput) (*Animal, error)
	Update(input *models.AnimalUpdateInput) (*Animal, error)
	Remove(id int64) error
	AddType(animal int64, typeId int64) error
	ChangeType(animal int64, oldType int64, newType int64) error
	RemoveType(animal int64, typeId int64) error
	Locations(input *models.AnimalVisitedLocationListInput) (*models.AnimalVisitedLocation, error)
	AddLocation(animal int64, location int64) error
	ChangeLocation(animal int64, visitedLocation int64, location int64) error
	RemoveLocation(animal int64, visitedLocation int64)
}

type Type interface {
	GetById(id int64) (*models.AnimalType, error)
	New(name string) error
	Modify(id int64, name string) error
	Remove(id int) error
}

type Location interface {
	GetById(id int64) (*models.Location, error)
	New(latitude float32, longitude float32) error
	Modify(id int64, latitude float32, longitude float32) error
	Remove(id int64)
}

type Service struct {
	Authorization
	Animal
	User
	Location
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: repos.Authorization,
		Animal:        repos.Animal,
		User:          repos.User,
		Location:      repos.Location,
	}
}

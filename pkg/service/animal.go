package service

import (
	"animal-chipization/models"
	"animal-chipization/pkg/repository"
)

type AnimalService struct {
	repo repository.Animal
}

func NewAnimalService(repo repository.Animal) *AnimalService {
	return &AnimalService{repo: repo}
}

func (s *AnimalService) GetById(id int64) (*models.Animal, error) {
	return s.repo.GetById(id)
}

func (s *AnimalService) Search(input *models.AnimalSearchInput) (*models.Animal, error) {
	return s.repo.Search(input)
}

func (s *AnimalService) New(input *models.AnimalNewInput) (*models.Animal, error) {
	return s.repo.New(input)
}

func (s *AnimalService) Update(input *models.AnimalUpdateInput) (*models.Animal, error) {
	return s.repo.Update(input)
}

func (s *AnimalService) Remove(id int64) error {
	return s.repo.Remove(id)
}

func (s *AnimalService) AddType(animal int64, typeId int64) error {
	return s.repo.AddType(animal, typeId)
}

func (s *AnimalService) ChangeType(animal int64, oldType int64, newType int64) error {
	return s.repo.ChangeType(animal, oldType, newType)
}

func (s *AnimalService) RemoveType(animal int64, typeId int64) error {
	return s.repo.RemoveType(animal, typeId)
}

func (s *AnimalService) Locations(input *models.AnimalVisitedLocationListInput) (*models.AnimalVisitedLocation, error) {
	return s.repo.Locations(input)
}

func (s *AnimalService) AddLocation(animal int64, location int64) error {
	return s.repo.AddLocation(animal, location)
}

func (s *AnimalService) ChangeLocation(animal int64, visitedLocation int64, location int64) error {
	return s.repo.ChangeLocation(animal, visitedLocation, location)
}

func (s *AnimalService) RemoveLocation(animal int64, visitedLocation int64) error {
	return s.repo.RemoveLocation(animal, visitedLocation)
}

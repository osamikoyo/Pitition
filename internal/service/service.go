package service

import (
	"github.com/osamikoyo/pitition/internal/data"
	"github.com/osamikoyo/pitition/internal/data/models"
)

type Pititioner interface {
	Add(pitition models.Pitition) error
	Get(name string) ([]models.Pitition, error)
	Like(name string) error
}
type Service struct {
	Storage data.Storage
}

func New() Pititioner {
	return Service{
		Storage: data.New(),
	}
}

func (s Service) Add(pitition models.Pitition) error {
	return s.Storage.Add(pitition)
}

func (s Service) Get(name string) ([]models.Pitition, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) Like(name string) error {
	//TODO implement me
	panic("implement me")
}

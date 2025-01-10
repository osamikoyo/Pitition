package data

import (
	"github.com/osamikoyo/pitition/internal/data/models"
	"github.com/osamikoyo/pitition/pkg/loger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Storage interface {
	Add(pitition models.Pitition) error
	Get(name string) (models.Pitition, error)
	UpdateLike(name string) error
}

type Database struct {
	*gorm.DB
}

func New() Database {
	db, err := gorm.Open(sqlite.Open("storage/main.db"))
	if err != nil {
		loger.New().Error().Err(err)
	}

	return Database{db}
}

func (d Database) Add(pitition models.Pitition) error {
	//TODO implement me
	panic("implement me")
}

func (d Database) Get(name string) (models.Pitition, error) {
	//TODO implement me
	panic("implement me")
}

func (d Database) UpdateLike(name string) error {
	//TODO implement me
	panic("implement me")
}

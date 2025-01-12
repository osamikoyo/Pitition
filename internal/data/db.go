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
	return d.Create(&pitition).Error
}

func (d Database) Get(name string) (models.Pitition, error) {
	var pit models.Pitition
	if err := d.Where(&models.Pitition{
		Title: name,
	}).Find(&pit).Error; err != nil {
		return models.Pitition{}, err
	}

	return pit, nil
}

func (d Database) UpdateLike(name string) error {
	var pit models.Pitition
	if err := d.Where(&models.Pitition{
		Title: name,
	}).Find(&pit).Error; err != nil {
		return err
	}

	return d.Model(&pit).Update("title", pit.Likes+1).Error
}

package models

type Pitition struct {
	ID      uint64 `gorm:"primaryKey"`
	Title   string
	Content string
	Likes   uint64
}
